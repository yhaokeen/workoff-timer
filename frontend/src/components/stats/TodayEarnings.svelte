<script lang="ts">
    import {onMount, onDestroy} from 'svelte';

    export let monthlySalary: number = 10000;
    export let workStartHour: number = 9;
    export let workEndHour: number = 18;

    let earnings = "0.000"
    let timer: number;

    function calculate() {
        const now = new Date();
        const dailySalary = monthlySalary / 22;
        const workStart = new Date(now.getFullYear(), now.getMonth(), now.getDate(), workStartHour, 0, 0);
        const workEnd = new Date(now.getFullYear(), now.getMonth(), now.getDate(), workEndHour, 0, 0);

        if (now >= workStart && now < workEnd) {
            const worked = now.getTime() - workStart.getTime();
            const totalWorked = workEnd.getTime() - workStart.getTime();
            const earnings = (dailySalary * worked / totalWorked).toFixed(3);
        } else if (now >= workEnd) {
            earnings = dailySalary.toFixed(3);
        } else {
            earnings = "0.000";
        }
    }

    onMount(() => {
        calculate();
        timer = window.setInterval(calculate, 1000);
    });

    onDestroy(() => {
        if (timer) {
            window.clearInterval(timer);
        }
    });
</script>

<div class="stat-item">
    <div class="label">今天赚了</div>
    <div class="value">{earnings}<span class="unit">¥</span></div>
</div>

<style>
    .stat-item { text-align: center; padding: 0 10px; }
    .label { font-size: 12px; color: #888; margin-bottom: 4px; }
    .value { font-size: 18px; font-weight: bold; color: #333; }
    .unit { font-size: 12px; font-weight: normal; color: #666; margin-left: 2px; }
  </style>