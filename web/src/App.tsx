import { createSignal, onMount, Show } from "solid-js";
import { LoggedInView } from "./pages/LoggedInView";
import Login from "./pages/Login";

export const App = (props: any) => {
  const [user, setUser] = createSignal(null);

  onMount(async () => {
    try {
      const resp = await fetch("/api/me", {
        method: "GET",
        credentials: "include", // so cookies are sent
      });
      if (resp.ok) {
        const data = await resp.json();
        setUser(data);
      } else {
        // Not logged in
        setUser(null);
      }
    } catch (error) {
      console.error(error);
      setUser(null);
    }
  });
  return (
    <div>
      <Show when={user()} fallback={<Login />}>
        <LoggedInView user={user()} {...props} />
      </Show>
    </div>
  );
};
