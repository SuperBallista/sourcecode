<script>
  import LginForm from "./LginForm.svelte";
  import Register from "./Register.svelte";
  import Findpw from "./Findpw.svelte";
  import Myinfo from "./Myinfo.svelte";
  import Changepw from "./Changepw.svelte";
  import Changeemail from "./Changeemail.svelte";
  import Newrecord from "./Newrecord.svelte";
  import Recordok from "./Recordok.svelte";
  import Newevent from "./Newevent.svelte";

  import { clickformopen, formopen, key } from "./store.js";

  function updatekey() {
    key.update((n) => n + 1);
  }
  function closeform() {
    clickformopen(null);
    updatekey();
  }
</script>

{#if $formopen}
  <div class="form">
    <div class="form-title">
      <div class="outbutton" on:click={() => closeform()}>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="size-6"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M6 18 18 6M6 6l12 12"
          />
        </svg>
      </div>
      <!-- 단순히 함수명을 넣어줍니다 -->
    </div>
    {#if $formopen === "login"}
      <!-- 비교 연산자 ===를 사용하는 것이 좋습니다 -->
      <LginForm />
    {:else if $formopen === "register"}
      <Register />
    {:else if $formopen === "findpw"}
      <Findpw />
    {:else if $formopen === "myinfo"}
      <Myinfo />
    {:else if $formopen === "changepw"}
      <Changepw />
    {:else if $formopen === "changeemail"}
      <Changeemail />
    {:else if $formopen === "newrecord"}
      <Newrecord />
    {:else if $formopen === "recordok"}
      <Recordok />
    {:else if $formopen === "newevent"}
      <Newevent />
    {/if}
  </div>
{/if}

<style>
  .form {
    margin: auto;
    width: 360px;
    background-color: #444;
    padding: 10px;
    border-radius: 15px;
    border: white 2px solid;
  }
  .form-title {
    display: flex;
    justify-content: flex-end; /* right 대신 flex-end를 사용합니다 */
  }
  .outbutton {
    padding: 3px 2px 1px 2px;
    width: 30px;
    height: 30px;
    cursor: pointer;
    border-radius: 5px;
    border: 2px solid white;
  }
</style>
