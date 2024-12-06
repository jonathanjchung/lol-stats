import * as React from "react";
import * as ReactDOM from "react-dom/client";
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import './index.css';
import Home from "./routes/home";
import ErrorPage from "./routes/error";
import Summoner from "./routes/summoner";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
    errorElement: <ErrorPage />, // does this need to be in every thingy
  },
  {
    path: "summoners/:summonerName/:tagLine",
    element: <Summoner />
  },
]);

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);