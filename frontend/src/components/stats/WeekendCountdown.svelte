<script lang="ts">
    import {onMount} from 'svelte';

    let days = 0;
    function calculate() {
        const dayOfWeek = new Date().getDay();
        if (dayOfWeek === 0) days = 5;
        else if (dayOfWeek === 6) days = 6;
        else days = 5 - dayOfWeek;
    }

    onMount(() => {
        calculate();
        const timer = window.setInterval(calculate, 1000 * 60 * 10);
        return () => window.clearInterval(timer);
    })
</script>

<div class="stat-item">
    <div class="label">周五</div>
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