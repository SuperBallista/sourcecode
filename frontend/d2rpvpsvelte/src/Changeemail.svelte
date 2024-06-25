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

  let pw = "";
  let newemail = "";

  onMount(async () => {
    await getCsrfToken();
  });

  async function changepw() {
    const token = get(csrfToken); // Svelte store에서 현재 값을 가져옴

    const emaildata = {
      nowpw: pw,
      newemail: newemail,
    };
    const checkjwt = get(jwtoken);
    if (checkjwt) {
      try {
        const endpoint = $mode
          ? "/process_changeemail_m"
          : "/process_changeemail";
        const response = await fetch(endpoint, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${checkjwt}`,
            "CSRF-Token": token,
          },
          body: JSON.stringify(emaildata), // Send the JSON data as the body of the request
        });
        alert("이메일 변경 완료");
        clickformopen(null);
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
  <input class="namewidth" type="password" bind:value={pw} />
  새 이메일
  <input class="namewidth" type="text" bind:value={newemail} />
  <button on:click={() => changepw()}>이메일 변경</button>
</div>
