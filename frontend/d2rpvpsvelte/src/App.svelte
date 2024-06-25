<script>
  import { Router, Route } from "svelte-routing";

  import Babapk from "./Babapk.svelte";
  import Notice from "./babapk/Notice.svelte";
  import Rank from "./babapk/Rank.svelte";
  import Record from "./babapk/Record.svelte";
  import Event from "./babapk/Event.svelte";
  import Calculater from "./babapk/Calculater.svelte";
  import Fame from "./babapk/Fame.svelte";
  import Winner from "./babapk/Winner.svelte";
  import Menu from "./Menu.svelte";
  import Visio from "./babapk/Visio.svelte";
  import Pscore from "./babapk/Pscore.svelte";
  import Stat from "./babapk/Stat.svelte";
  import Simul from "./babapk/Simul.svelte";
  import Ring from "./babapk/Ring.svelte";
  import Admin from "./babapk/Admin.svelte";
  import Oldrecord from "./babapk/Oldrecord.svelte";

  import Mpk from "./Mpk.svelte";
  import Mnotice from "./mpk/Notice.svelte";
  import Mrank from "./mpk/Rank.svelte";
  import Mevent from "./mpk/Event.svelte";
  import Mcalculater from "./mpk/Calculater.svelte";
  import Mwinner from "./mpk/Winner.svelte";
  import Mpscore from "./mpk/Pscore.svelte";
  import Msimul from "./mpk/Simul.svelte";
  import Mring from "./mpk/Ring.svelte";
  import Madmin from "./mpk/Admin.svelte";
  import Moldrecord from "./mpk/Oldrecord.svelte";

  // Store를 사용하여 컴포넌트 간에 공유 상태를 관리할 수 있습니다.

  import { writable } from "svelte/store";

  // 현재 URL을 가져와서 `mpk`가 포함되어 있는지 확인합니다.
  const initialValue = window.location.href.includes("mpk");

  // 초기값으로 스토어를 생성합니다.
  export const mode = writable(initialValue);

  // URL이 변경될 때 스토어 값을 업데이트할 수 있습니다. (선택 사항)
  function checkUrl() {
    mode.set(window.location.href.includes("mpk"));
  }

  // URL 변경을 감지하려면 이벤트 리스너를 추가할 수 있습니다.
  window.addEventListener("popstate", checkUrl);
  window.addEventListener("pushstate", checkUrl);
</script>

<Router>
  <Menu />

  <!-- 각 라우트에 key 속성을 추가하여 URL이 변경될 때마다 컴포넌트를 재렌더링 -->
  <Route path="/babapk" component={Babapk} />
  <Route path="/babapk/notice" component={Notice} />
  <Route path="/babapk/rank" component={Rank} />
  <Route path="/babapk/record" component={Record} />
  <Route path="/babapk/event" component={Event} />
  <Route path="/babapk/calculater" component={Calculater} />
  <Route path="/babapk/visio" component={Visio} />
  <Route path="/babapk/pscore" component={Pscore} />
  <Route path="/babapk/stat" component={Stat} />
  <Route path="/babapk/simul" component={Simul} />
  <Route path="/babapk/ring" component={Ring} />
  <Route path="/babapk/fame" component={Fame} />
  <Route path="/babapk/winner" component={Winner} />
  <Route path="/babapk/admin" component={Admin} />
  <Route path="/babapk/oldrecord" component={Oldrecord} />
  <Route path="/mpk" component={Mpk} />
  <Route path="/mpk/notice" component={Mnotice} />
  <Route path="/mpk/rank" component={Mrank} />
  <Route path="/mpk/event" component={Mevent} />
  <Route path="/mpk/calculater" component={Mcalculater} />
  <Route path="/mpk/pscore" component={Mpscore} />
  <Route path="/mpk/simul" component={Msimul} />
  <Route path="/mpk/ring" component={Mring} />
  <Route path="/mpk/winner" component={Mwinner} />
  <Route path="/mpk/admin" component={Madmin} />
  <Route path="/mpk/oldrecord" component={Moldrecord} />
</Router>
