(() => {
  "use static";

  let content = null;

  function getCookies(cname) {
    let name = cname + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let ca = decodedCookie.split(";");
    for (let i = 0; i < ca.length; i++) {
      let c = ca[i];
      while (c.charAt(0) == " ") {
        c = c.substring(1);
      }
      if (c.indexOf(name) == 0) {
        return c.substring(name.length, c.length);
      }
    }
    return "";
  }

  async function getCurrentPlaying() {
    return fetch("/api/playing", {
      method: "GET",
    })
      .then((res) => res.json())
      .then((res) => {
        return res;
      });
  }

  function toMinutes(ms) {
    let min = Math.floor((ms / 1000 / 60) << 0);
    let sec = Math.floor((ms / 1000) % 60);

    return `${min > 10 ? min : "0" + min}:${sec > 10 ? sec : "0" + sec}`;
  }

  function renderPlaying(data) {
    content.innerHTML = `
      <div class="flex items-center justify-between">
        <p class="text-sm">Now playing...</p>
        <button id="update" class="text-zinc-300 text-xs font-medium hover:text-zinc-300/900">Update</button>
      </div>
      <div class="flex p-3 border border-zinc-900 gap-3 rounded">
        <div class="h-20 w-20 bg-zinc-900 rounded overflow-hidden">
          <img loading="lazy" src="${data.item.album.images[1].url}" class="w-full h-full" alt="${data.item.album.name}" />
        </div>
        <div class="space-y-2">
          <h3 class="text-lg font-medium leading-tight">${data.item.name}</h3>
          <div>
            <p class="text-sm">${data.item.artists[0].name}</p>
            <p class="text-sm">${toMinutes(data.item.duration_ms)}</p>
          </div>
        </div>
      </div>
      <div class="flex items-center justify-between gap-3 w-full">
        <button id="prev" type="button" class="text-sm w-full font-medium py-1.5 px-5 rounded bg-white text-zinc-950 hover:bg-white/90">Prev</button>
        <button
          id="copylink"
          type="button"
          data-url="${data.item.external_urls.spotify}"
          class="bg-white hover:bg-white/90 py-1.5 px-5 justify-center items-center font-medium text-sm w-full flex text-zinc-950 rounded text-nowrap"
        >
          Copy Link
        </button>
        <button id="next" type="button" class="text-sm w-full font-medium py-1.5 px-5 rounded bg-white text-zinc-950 hover:bg-white/90">Next</button>
      </div>`;

    const copylink = document.getElementById("copylink");
    const update = document.getElementById("update");

    copylink.addEventListener("click", () => {
      navigator.clipboard.writeText(data.item.external_urls.spotify);
      copylink.textContent = "Link copied!";
      setTimeout(() => {
        copylink.textContent = "Copy Link";
      }, 1000);
    });

    update.addEventListener("click", () => {
      update.textContent = "Updating...";
      getPlaying().finally(() => {
        update.textContent = "Update";
      });
    });
  }

  let getPlaying = () =>
    getCurrentPlaying()
      .then((res) => {
        console.log(res);
        if (res.data) {
          renderPlaying(res.data);
        } else {
          content.innerHTML = `
             <div
                  class="flex p-3 border border-zinc-900 gap-3 rounded items-center justify-center"
              >
              <div>
                <h3 class="font-medium">Not Playing</h3>
              </div>
            </div>`;
        }
      })
      .catch((err) => {
        console.error(err);
        content.innerHTML = `
           <div
                class="flex p-3 border border-zinc-900 gap-3 rounded items-center justify-center"
            >
            <div>
              <h3 class="font-medium">Someting Wrong</h3>
            </div>
          </div>`;
      });

  document.addEventListener("DOMContentLoaded", () => {
    const login = document.getElementById("login");
    content = document.getElementById("content");

    if (getCookies("AccessToken")) {
      login.remove();
      content.classList.remove("hidden");

      getPlaying();

      setInterval(() => {
        getPlaying();
      }, 1000 * 60);
    }
  });
})();
