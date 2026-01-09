<script lang="ts">
    import {onMount} from 'svelte';
    import {GetNextFestival} from '../../../wailsjs/go/main/App';

    let festivalName = "";
    let daysLeft = 0;

    async function loadFestival() {
        const info = await GetNextFestival();
        festivalName = info.name;
        daysLeft = info.days;

    }

    onMount(() => {
        loadFestival();
        const timer = window.setInterval(loadFestival, 1000 * 60 * 60);
        return () => window.clearInterval(timer);
    })
</script>

<div class="stat-item">
    <div class="label">{festivalName}</div>
    <div class="value">{daysLeft}<span class="unit">天</span></div>
</div>

<style>
    .stat-item { text-align: center; padding: 0 10px; }
    .label { font-size: 12px; color: #888; margin-bottom: 4px; }
    .value { font-size: 18px; font-weight: bold; color: #333; }
    .unit { font-size: 12px; font-weight: normal; color: #666; margin-left: 2px; }
</style>