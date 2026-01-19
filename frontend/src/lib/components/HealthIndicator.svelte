<script lang="ts">
  import { onMount } from 'svelte';

  let status = 'checking';
  let uptime = '';
  let intervalId: ReturnType<typeof setInterval>;


  async function checkHealth() {
    try {
      const res = await fetch('http://127.0.0.1:3000/api/v1/health');
      if (res.ok) {
        const data = await res.json();
        status = 'ok';
        uptime = data.uptime;
      } else {
        status = 'error';
        uptime = '';
      }
    } catch (err) {
      status = 'error';
      uptime = '';
    }
  }

  onMount(() => {
    checkHealth();
    intervalId = setInterval(checkHealth, 5000);
    return () => clearInterval(intervalId);
  });
</script>

<div
  style="position: fixed; bottom: 16px; right: 16px; display: flex; align-items: center; gap: 8px; font-size: 0.85rem;"
  title={uptime ? `Uptime: ${uptime}` : 'Сервис недоступен'}
>
  <span
    style="
      width: 10px;
      height: 10px;
      border-radius: 50%;
      background: {
        status === 'ok' ? 'var(--status-ok)' :
        status === 'error' ? 'var(--danger)' : '#9ca3af'
      };
    "
  ></span>
  <span>API</span>
</div>