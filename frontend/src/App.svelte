<script lang="ts">
  import CountdownTimer from "./components/CountdownTimer.svelte";
  import PaydayCountdown from "./components/stats/PaydayCountdown.svelte";
  import WeekendCountdown from "./components/stats/WeekendCountdown.svelte";
  import TodayEarnings from "./components/stats/TodayEarnings.svelte";
  import HolidayCountdown from "./components/stats/HolidayCountdown.svelte";
</script>

<main>
  <div class="card" style="--wails-draggable:drag">
    <div class="content">
      <CountdownTimer offWorkHour={19} offWorkMinute={0} title="下班还有" />
      <div class="stats">
        <PaydayCountdown payday={10} />
        <WeekendCountdown />
        <HolidayCountdown />
        <TodayEarnings monthlySalary={10000} />
      </div>
    </div>
  </div>
</main>

<style>
  main {
    width: 100vw;
    height: 100vh;
    margin: 0;
    padding: 0;
    background: transparent;
  }

  .card {
    width: 100%;
    height: 100%;
    /* 通过这里控制透明度，0.85 = 85%不透明 */
    background: rgba(255, 255, 255, 0.6);
    border-radius: 16px;
    padding: 20px 30px;
    box-sizing: border-box;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
    color: #333;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
  .stats {
    display: flex;
    flex-direction: row;
    gap: 20px;
    margin-top: 15px;
    padding-top: 15px;
    border-top: 1px solid rgba(0, 0, 0, 0.1);
  }
</style>
