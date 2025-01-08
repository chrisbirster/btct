import { createResource, For, Match, Show, Switch } from "solid-js";

import { css } from "@linaria/core";
import { fetchTasks } from "../utils";

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
  ID: number;
  Description: string;
  Complete: boolean;
};

export default function Tasks() {
  const [tasks, { mutate }] = createResource<Task[]>(fetchTasks);

  const addTask = async (event: Event) => {
    event.preventDefault();

    const form = event.currentTarget as HTMLFormElement;
    const formData = new FormData(form);
    const description = formData.get("description") as string;

    try {
      const resp = await fetch("/api/tasks/create", {
        method: "POST",
        body: new URLSearchParams({ description }),
        credentials: "include",
      });
      if (!resp.ok) {
        const errText = await resp.text();
        throw new Error(errText);
      }
      const newTask = await resp.json();
      mutate([...(tasks() ?? []), newTask]);
      form.reset();
    } catch (error) {
      console.error("Error creating task:", error);
    }
  };

  return (
    <div class={container}>
      <h1>Tasks</h1>
      <Show when={tasks.loading}>
        <p>Loading...</p>
      </Show>
      <Switch>
        <Match when={tasks.error}>
          <span>Error: {tasks.error?.message}</span>
        </Match>
        <Match when={tasks()}>
          <Show when={tasks()} keyed>
            {(tasks) => {
              return (
                <ul>
                  <For each={tasks} fallback={<p>Loading...</p>}>
                    {(task) => <p>{task.Description}</p>}
                  </For>
                </ul>
              );
            }}
          </Show>
        </Match>
      </Switch>
      <form onSubmit={addTask}>
        <label for="description">Task Description</label>
        <input
          id="description"
          name="description"
          type="text"
          placeholder="What is this task about?"
        />
        <button class={button} type="submit">
          Add Task
        </button>
      </form>
    </div>
  );
}
