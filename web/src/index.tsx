/* @refresh reload */
import { lazy } from "solid-js";
import { render } from "solid-js/web";
import { Router, Route } from "@solidjs/router";
import { App } from "./App";

const Home = lazy(() => import("./pages/Home"));
const Tasks = lazy(() => import("./pages/Tasks"));
const Task = lazy(() => import("./pages/Task"));
const NotFound = lazy(() => import("./pages/NotFound"));

const root = document.getElementById("root");

if (!root) {
  throw new Error("element not found");
} else {
  render(
    () => (
      <Router root={App}>
        <Route path="/" component={Home} />
        <Route path="/tasks" component={Tasks} />
        <Route path="/tasks/:id" component={Task} />
        <Route path="*paramName" component={NotFound} />
      </Router>
    ),
    root,
  );
}
