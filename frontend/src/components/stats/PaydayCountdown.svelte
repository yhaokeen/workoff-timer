<script lang="ts">
    import {onMount} from 'svelte';

    export let payday: number = 15;
    let days = 0;

    function calculate() {
        const now = new Date();
        const currentDay = now.getDate();

        if (currentDay <= payday) {
            days = payday - currentDay;
        } else {
            const nextMonth = new Date(now.getFullYear(), now.getMonth() + 1, payday);
            days = nextMonth.getDate() - currentDay;
        }
    }

    onMount(() => {
        calculate();
        const timer = window.setInterval(calculate, 1000 * 60 * 10)  // 10min更新一次就够了
        return () => window.clearInterval(timer);
    });
</script>

<div class="stat-item">
    <div class="label">发薪</div>
    <div class="value">{days}<span class="unit">天</span></div>
</div>

<style>
    .stat-item {
        text-align: center;
        padding: 0 10px;
    }
    .label {font-size: 12px; color: #888; margin-bottom: 4px}
    .value {font-size: 18px; font-weight: bold; color: #333}
    .unit {font-size: 12px; font-weight: normal; color: #666; margin-left: 4px}
</style>