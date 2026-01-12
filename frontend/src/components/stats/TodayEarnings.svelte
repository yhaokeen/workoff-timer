<script lang="ts">
    import {onMount, onDestroy} from 'svelte';
    import StatItem from './StatItem.svelte';

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
        timer = window.setInterval(calculate, 100);
    });

    onDestroy(() => {
        if (timer) {
            window.clearInterval(timer);
        }
    });
</script>


<StatItem label="今天赚了" value={earnings} unit="¥" />
