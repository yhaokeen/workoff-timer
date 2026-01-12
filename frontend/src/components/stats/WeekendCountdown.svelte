<script lang="ts">
    import {onMount} from 'svelte';
    import StatItem from './StatItem.svelte';

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

<StatItem label="周五" value={days} unit="天" />
