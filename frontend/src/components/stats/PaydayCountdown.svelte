<script lang="ts">
    import {onMount} from 'svelte';
    import StatItem from './StatItem.svelte';

    export let payday: number = 15;
    let days = 0;

    function calculate() {
        const now = new Date();
        const currentDay = now.getDate();

        if (currentDay <= payday) {
            days = payday - currentDay;
        } else {
            const nextMonth = new Date(now.getFullYear(), now.getMonth() + 1, payday);
            days = Math.ceil((nextMonth.getTime() - now.getTime()) / (1000 * 60 * 60 * 24));
        }
    }

    onMount(() => {
        calculate();
        const timer = window.setInterval(calculate, 1000 * 60 * 10)  // 10min更新一次就够了
        return () => window.clearInterval(timer);
    });
</script>

<StatItem label="发薪" value={days} unit="天" />