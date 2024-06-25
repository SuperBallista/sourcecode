<script>
  import {
    clickformopen,
    mode,
    nickname,
    getCsrfToken,
    csrfToken,
    jwtoken,
  } from "./store.js";
  import { onMount } from "svelte";
  import { get } from "svelte/store";

  let loginnickname = "";
  let loginpw = "";

  onMount(async () => {
    await getCsrfToken();
  });

  // JWT 디코딩을 위한 함수 (Base64Url 디코딩)
  function decodeJwt(token) {
    const base64Url = token.split(".")[1];
    const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
    const jsonPayload = decodeURIComponent(
      atob(base64)
        .split("")
        .map((c) => {
          return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
        })
        .join("")
    );

    return JSON.parse(jsonPayload);
  }

  // 로그인 프로세스 함수
  async function formprocess() {
    let logindata = { nickname: loginnickname, password: loginpw };
    const token = get(csrfToken); // Svelte store에서 현재 값을 가져옴
    try {
      const endpoint = $mode ? "/process_login_m" : "/process_login";
      const response = await fetch(endpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "CSRF-Token": token,
        },
        credentials: "include", // 쿠키 포함
        body: JSON.stringify(logindata),
      });

      if (response.status === 200) {
        const data = await response.json();
        jwtoken.set(data.token); // 서버에서 받은 JWT 토큰
        const decoded = decodeJwt($jwtoken);
        nickname.set(decoded.username);
        clickformopen(null);
      } else {
        alert("로그인 실패: " + (await response.text())); // 에러 메시지 출력
      }
    } catch (error) {
      alert("서버 오류입니다");
    }
  }
</script>

<div class="left">
  닉네임
  <input class="namewidth" type="text" bind:value={loginnickname} />
  암호
  <input class="namewidth" type="password" bind:value={loginpw} />
  <button class="posit1" on:click={() => formprocess()}>로그인</button>
  <button class="posit2" on:click={() => clickformopen("findpw")}
    >암호찾기</button
  >
</div>

<style>
  .posit1 {
    position: relative;
    top: -35px;
    left: 280px;
  }
  .posit2 {
    position: relative;
    left: -70px;
  }
</style>
