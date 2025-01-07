import { A } from "@solidjs/router";

export const LoggedInView = (props: any) => {
  console.log(props);
  return (
    <div>
      <p>Welcome back, {props.user.email}!</p>
      <nav>
        <A href="/">Home</A>
        <A href="/login">Login</A>
        <A href="/tasks">Tasks</A>
      </nav>
      {props.children}
    </div>
  );
};
