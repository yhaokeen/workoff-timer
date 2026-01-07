<script lang="ts">
  import { onMount, onDestroy } from 'svelte';

    // Props - 从父组件接收的参数
  export let offWorkHour: number = 18;
  export let offWorkMinute: number = 0;
  export let title: string = "下班还有";
  // 组件内部状态
  let countdown = "00:00:00";
  let timer: number;

  function updateCountdown() {
    const now = new Date();
    const target = new Date();
    target.setHours(offWorkHour, offWorkMinute, 0, 0);
    let diff = target.getTime() - now.getTime();
    if (diff < 0) {
      countdown = "00:00:00";
      return;
    }

    const hours = Math.floor(diff / (1000 * 60 * 60));
    diff -= hours * (1000 * 60 * 60);
    const minutes = Math.floor(diff / (1000 * 60));
    diff -= minutes * (1000 * 60);
    const seconds = Math.floor(diff / 1000);
    
    countdown = [hours, minutes, seconds].map(num => num.toString().padStart(2, '0')).join(':')
  }

  onMount(() => {
    updateCountdown();
    timer = window.setInterval(updateCountdown, 1000);
  });

  onDestroy(() => {
    if (timer) {
      window.clearInterval(timer);
    }
  });
</script>

<div class="countdown-container">
  <div class="header">{title}</div>
  <div class="countdown">{countdown}</div>
</div>

<style>
    .countdown-container {
      text-align: center;
    }
    
    .header {
      font-size: 14px;
      color: #666;
      margin-bottom: 8px;
    }
    
    .countdown {
      font-size: 48px;
      font-weight: bold;
      color: #333;
      font-family: "Consolas", monospace;
    }
  </style>