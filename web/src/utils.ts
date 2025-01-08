export const fetchTasks = async () => {
  try {
    const resp = await fetch("/api/tasks", {
      method: "GET",
      credentials: "include",
    });
    if (!resp.ok) {
      const errText = await resp.text();
      throw new Error(errText);
    }
    const data = await resp.json();
    return data;
  } catch (error) {
    console.error("Error creating task:", error);
  }
};
