<script lang="ts">
    import {onMount} from 'svelte';
    import {GetNextFestival} from '../../../wailsjs/go/main/App';
    import StatItem from './StatItem.svelte';

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

<StatItem label={festivalName} value={daysLeft} unit="天" />
