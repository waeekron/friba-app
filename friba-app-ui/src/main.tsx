import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import {
    createBrowserRouter,
    RouterProvider,
} from "react-router-dom";
import JoinScorecard from "./routes/JoinScorecard.tsx";
import CreateScorecard from "./routes/CreateScorecard.tsx";
import Scorecard from "./routes/Scorecard.tsx";
const router = createBrowserRouter([
    {
        path:"/",
        element: <App />, children: [
            {path: "uusi-peli", element: <CreateScorecard/>},
            {path: "liity-peliin", element: <JoinScorecard/>}
        ]
    },
    {
        path: "tuloskortti/:koodi",
        element: <Scorecard />
    }
])
ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
      <header>
          <div className="drawer ">
              <input id="my-drawer-3" type="checkbox" className="drawer-toggle" />
              <div className="drawer-content z-10 flex flex-col ">
                  {/* Navbar */}
                  <div className="w-full navbar  p-0 min-h-12 ">
                      <div className="flex-1 lg:hidden">
                          <label htmlFor="my-drawer-3" aria-label="open sidebar" className="btn btn-square btn-ghost">
                              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" className="inline-block w-6 h-6 stroke-current"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4 6h16M4 12h16M4 18h16"></path></svg>
                          </label>
                      </div>
                      <div className="flex-none px-2 mx-2"><span className="rounded-full bg-accent font-extrabold font-mono p-2">FA</span></div>
                      <div className="flex-none hidden lg:block">
                          <ul className="menu menu-horizontal">
                              {/* Navbar menu content here */}
                              <li><a>Navbar Item 1</a></li>
                              <li><a>Navbar Item 2</a></li>
                          </ul>
                      </div>
                  </div>
                  <RouterProvider router={router} ></RouterProvider>
              </div>
              <div className="drawer-side z-20">
                  <label htmlFor="my-drawer-3" aria-label="close sidebar" className="drawer-overlay"></label>
                  <ul className="menu p-4 w-80 min-h-full bg-base-200">
                      {/* Sidebar content here */}
                      <li><a>Sidebar Item 1</a></li>
                      <li><a>Sidebar kissa 2</a></li>
                  </ul>
              </div>
          </div>
      </header>


  </React.StrictMode>,
)
