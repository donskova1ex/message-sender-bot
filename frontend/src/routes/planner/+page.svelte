<div class="planner-card">
  <div class="columns">
    <div class="column">
      <div class="planner-header">
        <h1>Форма создания</h1>
      </div>
      <div class="form-content">
        <div class="form-row">
          <div class="form-col">
            <div class="form-group">
              <label for="planned-date">Дата отправки</label>
              <input type="datetime-local" class="form-control" id="planned-date" name="planned-date" required />
            </div>
          </div>
          <div class="form-col">
            <div class="form-group">
              <label for="message-type">Тип сообщения</label>
              <select class="form-control" id="message-type" name="message-type">
                <option value="">— Выберите тип —</option>
                <option value="1">Напоминание</option>
                <option value="2">Рассылка</option>
                <option value="3">Опрос</option>
                <option value="4">Приветствие</option>
              </select>
            </div>
          </div>
        </div>
        <div class="form-group full-height">
          <label for="message-text">Текст сообщения</label>
          <textarea
            id="message-text"
            name="message-text"
            class="form-control message-textarea"
            placeholder="Введите текст сообщения..."
            required
          ></textarea>
        </div>
        <button type="submit" class="btn">Запланировать</button>
      </div>
    </div>
    <div class="column">
      <div class="planner-header">
        <h1>Список запланированных</h1>
      </div>
    </div>
  </div>
</div>

<style>
  .planner-card {
    width: 100%;
    height: 700px;
    background: var(--bg-card);
    border: 1px solid var(--border);
    border-radius: var(--radius);
    padding: 32px;
    box-shadow: var(--shadow);
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .columns {
    display: flex;
    gap: 24px;
    width: 100%;
    height: 100%;
  }

  .column {
    flex: 1;
    background: var(--bg-card);
    border: 1px solid var(--border);
    border-radius: var(--radius);
    padding: 24px;
    box-shadow: var(--shadow);
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 100%;
  }

  .form-content {
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 100%;
    gap: 16px;
  }

  .form-row {
    display: flex;
    gap: 16px;
    width: 100%;
  }

  .form-col {
    flex: 1;
  }

  .full-height {
    flex: 1;
    display: flex;
    flex-direction: column;
    min-height: 0;
  }

  .message-textarea {
    flex: 1;
    min-height: 0;
    resize: none;
    overflow-y: auto;
    padding: 12px 16px;
    border: 1px solid var(--border);
    border-radius: 8px;
    background: var(--bg-card);
    color: var(--text);
    font-size: 1rem;
    font-family: inherit;
    box-sizing: border-box;
  }

  .planner-header {
    text-align: center;
    margin-bottom: 24px;
  }

  .planner-header h1 {
    font-weight: 700;
    font-size: 1.5rem;
    margin-bottom: 8px;
  }

  .btn {
    width: 50%;
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

  .form-control {
    width: 100%;
    padding: 12px 16px;
    border: 1px solid var(--border);
    border-radius: 8px;
    background: var(--bg-card);
    color: var(--text);
    font-size: 1rem;
    box-sizing: border-box;
    font-family: inherit;
  }

  input.form-control,
  select.form-control {
    height: 44px;
  }

  .form-group {
    margin-bottom: 0;
    width: 100%;
    display: flex;
    flex-direction: column;
  }

  .form-group label {
    display: block;
    margin-bottom: 6px;
    font-weight: 500;
    font-size: 0.95rem;
    text-align: center;
  }
</style>