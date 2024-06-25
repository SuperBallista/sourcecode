<script>
  import {
    csrfToken,
    getCsrfToken,
    jwtoken,
    mode,
    clickformopen,
  } from "./store.js";
  import { get } from "svelte/store";
  import { onMount } from "svelte";

  let oldpw = "";
  let newpw = "";
  let checkpw = "";

  onMount(async () => {
    await getCsrfToken();
  });

  async function changepw() {
    const token = get(csrfToken); // Svelte store에서 현재 값을 가져옴

    const pwdata = {
      nowpw: oldpw,
      newpw: newpw,
    };
    const checkjwt = get(jwtoken);
    if (checkjwt) {
      try {
        if (newpw !== checkpw) {
          alert("암호 확인이 일치하지 않습니다 다시 확인해주세요");
          return;
        } else {
          const endpoint = $mode ? "/process_changepw_m" : "/process_changepw";
          const response = await fetch(endpoint, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              Authorization: `Bearer ${checkjwt}`,
              "CSRF-Token": token,
            },
            body: JSON.stringify(pwdata), // Send the JSON data as the body of the request
          });
          if (response.ok) {
            alert("암호 변경 완료");
            clickformopen(null);
          } else {
            alert("에러 발생");
          }
        }
      } catch (error) {
        alert("오류 발생 :", error);
      }
    } else {
      alert("다시 로그인해주세요");
    }
  }
</script>

<div class="left">
  암호
  <input class="namewidth" type="password" bind:value={oldpw} />
  새 암호
  <input class="namewidth" type="password" bind:value={newpw} />
  암호 확인
  <input class="namewidth" type="password" bind:value={checkpw} />
  <button on:click={() => changepw()}>암호 변경</button>
</div>
