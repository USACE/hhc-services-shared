import { createRouteBundle } from "redux-bundler";
import Home from "../app-pages/home";
import NotFound from "../app-pages/404";

const base = import.meta.env.BASE_URL;

export default createRouteBundle({
  [`${base}`]: Home,
  [`${base}*`]: NotFound,
});
