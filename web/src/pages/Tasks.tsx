import { createSignal } from "solid-js";

import { css } from "@linaria/core";

const button = css`
  padding: 10px 20px;
  border-radius: 10px;
  color: #bada55;
`;

const container = css`
  margin: 40px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  font-size: 2rem;
`;

type Task = {
  id: number;
  description: string;
  complete: boolean;
};

export default function Tasks() {
  const [tasks, _] = createSignal<Task[]>();

  const addTask = (event: any) => {
    console.dir(event);
  };

  return (
    <div class={container}>
      <h1>Tasks</h1>
      <p>{JSON.stringify(tasks)}</p>
      <form method="post" action="/api/tasks/create">
        <label for="description">Task Description</label>
        <input
          id="description"
          name="description"
          type="text"
          placeholder="What is this task about?"
        />
        <button class={button} type="submit" onClick={addTask}>
          Add Task
        </button>
      </form>
    </div>
  );
}
