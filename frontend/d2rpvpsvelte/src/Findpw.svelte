<script>
  import { csrfToken, getCsrfToken, mode, clickformopen } from "./store.js";
  import { get } from "svelte/store";
  import { onMount } from "svelte";

  let findnickname = "";
  let findemail = "";

  onMount(async () => {
    await getCsrfToken();
  });

  async function findpw() {
    const token = get(csrfToken); // Svelte store에서 현재 값을 가져옴

    const findpwdata = {
      findpw_nickname: findnickname,
      findpw_email: findemail,
    };

    try {
      const endpoint = $mode ? "/process_emailpw_m" : "/process_emailpw";
      const response = await fetch(endpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "CSRF-Token": token,
        },
        credentials: "include", // 쿠키 포함
        body: JSON.stringify(findpwdata), // Send the JSON data as the body of the request
      });
      if (response.ok) {
        alert("이메일을 확인하세요");
        clickformopen(null);
      } else {
        alert("에러 발생");
      }
    } catch (error) {
      alert("오류 발생 :", error);
    }
  }
</script>

<div class="left">
  닉네임
  <input class="namewidth" type="text" bind:value={findnickname} />
  이메일
  <input class="namewidth" type="text" bind:value={findemail} />
  <button on:click={() => findpw()}>암호찾기</button>
</div>
