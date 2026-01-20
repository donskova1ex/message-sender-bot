<script>
    import { goto } from "$app/navigation";
    import HealthIndicator from "$lib/components/HealthIndicator.svelte";
    let email = '';
    let password = '';
    let loading = false;
    let error = false;

    async function handleSubmit() {
        if (loading) return;

        loading = true;
        error = false;

        try {
            const res = await fetch('http://localhost:3000/api/v1/auth/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({email, password})
            });
            
            if (res.ok) {
                await goto('/planner');
            } else {
                error = true;
            }
        } catch (err) {
            error = true;
        } finally {
            loading = false;
        }
    }
</script>

<div class="login-card">
    <div class="login-header">
        <h1>Панель планирования отправок</h1>
        <p>Войдите, чтобы начать работу</p>
    </div>
    {#if error}
        <div id="error" class="error-message">Неверный логин или пароль</div>
    {/if}
    
    <form on:submit|preventDefault={handleSubmit}>
        <div class="form-group">
            <label for="email">Email</label>
            <input
                type="email"
                id="email"
                bind:value={email}
                name="email"
                class="form-control"
                required
            />
        </div>
        <div class="form-group">
            <label for="password">Пароль</label>
            <input
                type="password"
                id="password"
                bind:value={password}
                name="password"
                class="form-control"
                required
            />
        </div>
        <button type="submit" class="btn" disabled={loading}>
            {loading ? 'Вход...': 'Войти'}
        </button>
    </form>
</div>
<HealthIndicator/>
<style>
    .login-card {
        width: 100%;
        max-width: 400px;
        background: var(--bg-card);
        border: 1px solid var(--border);
        border-radius: var(--radius);
        padding: 32px;
        box-shadow: var(--shadow);
    }
    .login-header {
        text-align: center;
        margin-bottom: 24px;
    }
    .login-header h1 {
        display: flex;
        justify-content: center;
        align-items: center;
        font-weight: 700;
        font-size: 1.5rem;
        margin-bottom: 8px;
    }
    .login-header p {
        display: flex;
        justify-content: center;
        align-items: center;
        color: var(--text-muted);
        font-size: 0.95rem;
    }
    .form-group {
        margin-bottom: 16px;
    }
    .form-group label {
        display: block;
        margin-bottom: 6px;
        font-weight: 500;
        font-size: 0.95rem;
    }
    .form-control {
        width: 100%;
        padding: 12px 16px;
        border: 1px solid var(--border);
        border-radius: 8px;
        background: var(--bg-card);
        color: var(--text);
        font-size: 1rem;
    }
    .form-control:focus {
        outline: none;
        border-color: var(--primary);
    }
    .btn {
        width: 100%;
        padding: 12px;
        background: var(--primary);
        color: white;
        border: none;
        border-radius: 8px;
        font-weight: 600;
        font-size: 1rem;
        cursor: pointer;
        transition: background 0.2s;
    }
    .btn:hover {
        background: var(--primary-hover);
    }
    .error-message {
        display: none;
        padding: 12px;
        background: rgba(239, 68, 68, 0.1);
        border: 1px solid var(--danger);
        border-radius: 8px;
        color: var(--danger);
        font-size: 0.9rem;
        margin-bottom: 16px;
    }
</style>
