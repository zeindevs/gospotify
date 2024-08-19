(() => {
  "use static";

  let content = null;

  let loveIcon = (fill = false) =>
    `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="${fill ? "#fff" : "none"}" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-heart h-4 w-4"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>`;

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

  async function currentPlaying() {
    return fetch("/api/playing", {
      method: "GET",
    })
      .then((res) => res.json())
      .then((res) => {
        return res;
      });
  }

  async function playPrev() {
    return fetch("/api/playing/prev", {
      method: "POST",
    })
      .then((res) => res.json())
      .then((res) => {
        return res;
      });
  }

  async function playNext() {
    return fetch("/api/playing/next", {
      method: "POST",
    })
      .then((res) => res.json())
      .then((res) => {
        return res;
      });
  }

  async function save(ids, is_saved) {
    return fetch("/api/track/save", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ids: [ids], is_saved }),
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
      <div class="flex p-3 border border-zinc-900 gap-3 rounded w-full">
        <div class="h-20 w-20 bg-zinc-900 rounded overflow-hidden">
          <img loading="lazy" src="${data.item.album.images[1].url}" class="w-full h-full" alt="${data.item.album.name}" />
        </div>
        <div class="space-y-2 flex-1">
          <h3 class="text-lg font-medium leading-tight">${data.item.name}</h3>
          <div>
            <p class="text-sm">${data.item.artists[0].name}</p>
            <p class="text-sm">${toMinutes(data.item.duration_ms)}</p>
          </div>
        </div>
        <div class="">
          <button id="saved">
            ${loveIcon(data.is_saved)}
          </button>
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

    const update = document.getElementById("update");
    const saved = document.getElementById("saved");
    const copylink = document.getElementById("copylink");
    const prev = document.getElementById("prev");
    const next = document.getElementById("next");

    update.addEventListener("click", () => {
      update.textContent = "Updating...";
      getPlaying().finally(() => {
        update.textContent = "Update";
      });
    });

    saved.addEventListener("click", () => {
      save(data.item.id, data.is_saved).then((res) => {
        console.log("save track:", res);
        if (res.data?.error) {
          alert(res.data?.error?.message);
        } else {
          getPlaying();
        }
      });
    });

    copylink.addEventListener("click", () => {
      navigator.clipboard.writeText(data.item.external_urls.spotify);
      copylink.textContent = "Link copied!";
      console.log("copied:", data.item.external_urls.spotify);
      setTimeout(() => {
        copylink.textContent = "Copy Link";
      }, 1000);
    });

    prev.addEventListener("click", () => {
      prev.textContent = "...";
      playPrev()
        .then((res) => {
          console.log("prev play:", res);
          if (res.data?.error) {
            alert(res.data?.error?.message);
          } else {
            getPlaying();
          }
        })
        .finally(() => {
          prev.textContent = "Prev";
        });
    });

    next.addEventListener("click", () => {
      next.textContent = "...";
      playNext()
        .then((res) => {
          console.log("next play:", res);
          if (res.data?.error) {
            alert(res.data?.error?.message);
          } else {
            getPlaying();
          }
        })
        .finally(() => {
          next.textContent = "Next";
        });
    });
  }

  async function getPlaying() {
    return currentPlaying()
      .then((res) => {
        console.log("current playing:", res);
        if (res.data) {
          try {
            renderPlaying(res.data);
          } catch (err) {
            content.innerHTML = `
             <div
                class="flex p-3 border border-zinc-900 gap-3 rounded items-center justify-center"
              >
              <div>
                <h3 class="font-medium">Someting Wrong</h3>
              </div>
            </div>`;
          }
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
  }

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
