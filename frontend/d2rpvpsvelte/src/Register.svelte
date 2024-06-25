<script>
  import { mode, csrfToken, clickformopen } from "./store";
  import { writable } from "svelte/store";

  let terms = writable(false);
  let privacy = writable(false);
  let reginickname = writable("");
  let regipw = writable("");
  let regipwcheck = writable("");
  let email = writable("");

  async function register() {
    // Check if terms and privacy are agreed
    if (!$terms || !$privacy) {
      alert("서비스 이용 약관과 개인정보 정책에 동의해야 합니다.");
      return;
    }

    // Check if password matches
    if ($regipw !== $regipwcheck) {
      alert("암호가 일치하지 않습니다.");
      return;
    }

    // Create the payload
    const payload = {
      nickname: $reginickname,
      password: $regipw,
      email: $email,
    };

    try {
      const endpoint = $mode ? "/process_regi_m" : "/process_regi";
      const response = await fetch(endpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "csrf-token": $csrfToken, // CSRF 토큰을 헤더에 포함
        },
        body: JSON.stringify(payload),
      });
      console.log($csrfToken);
      if (!response.ok) {
        throw new Error("회원가입에 실패했습니다.");
      }

      const result = await response.json();
      alert("회원가입이 성공적으로 완료되었습니다.");
      clickformopen(null);
    } catch (error) {
      alert(error.message);
    }
  }
</script>

<div class="left">
  {#if $mode}통합밀리{:else}정통바바{/if}
  <br />
  <a href="https://d2rpvp.org/site.html" target="_blank"
    ><div class="buttonLink">이용 약관</div></a
  >
  <label>
    <input type="checkbox" bind:checked={$terms} />서비스 이용 약관에
    동의합니다.</label
  >
  <a href="https://d2rpvp.org/privacy.html" target="_blank"
    ><div class="buttonLink">개인정보 정책</div></a
  >
  <label>
    <input type="checkbox" bind:checked={$privacy} />
    개인정보 정책에 동의합니다.
  </label>
  <br />

  닉네임
  <input class="namewidth" type="text" bind:value={$reginickname} />
  암호
  <input class="namewidth" type="password" bind:value={$regipw} />
  암호확인
  <input class="namewidth" type="password" bind:value={$regipwcheck} />
  이메일
  <input class="namewidth" type="text" bind:value={$email} />

  <button on:click={register}>등록하기</button>
</div>
