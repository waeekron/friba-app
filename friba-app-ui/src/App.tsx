import React from "react";
import {Link, Outlet} from "react-router-dom";

// function HomeActions() {
//     const [active, setActive] = useState(0)
//
//     return (
//         <>
//             <div role="tablist" className="tabs mx-2 mt-1 tabs-lifted h-[70svh] tabs-lg ">
//                 <input type="radio" name="my_tabs_2" role="tab" className="tab" aria-label="Tab 1"
//                        checked={active === 0} onChange={() => setActive(0)}/>
//                 <div role="tabpanel" className="tab-content bg-base-100 border-base-300 rounded-box ">
//                     <div className="h-[70svh]">
//                         <h1>Uusi peli</h1>
//                         <form action="">
//
//                         </form>
//                     </div>
//                 </div>
//                 <input type="radio" name="my_tabs_2" role="tab" className="tab" aria-label="Tab 1"
//                        checked={active === 0} onChange={() => setActive(1)}/>
//                 <div role="tabpanel" className="tab-content bg-base-100 border-base-300 rounded-box ">
//                     <div className="h-[70svh]"></div>
//                 </div>
//
//             </div>
//         </>
//
//     )
// }

function App() {

    return (
        <>

            <Outlet/>
            <nav className="btm-nav border-accent/10 border-t-2  px-2 gap-2 ">
                <Link to="liity-peliin" className="">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5"
                         stroke="currentColor" className="w-6 h-6">
                        <path strokeLinecap="round" strokeLinejoin="round"
                              d="M18 18.72a9.094 9.094 0 0 0 3.741-.479 3 3 0 0 0-4.682-2.72m.94 3.198.001.031c0 .225-.012.447-.037.666A11.944 11.944 0 0 1 12 21c-2.17 0-4.207-.576-5.963-1.584A6.062 6.062 0 0 1 6 18.719m12 0a5.971 5.971 0 0 0-.941-3.197m0 0A5.995 5.995 0 0 0 12 12.75a5.995 5.995 0 0 0-5.058 2.772m0 0a3 3 0 0 0-4.681 2.72 8.986 8.986 0 0 0 3.74.477m.94-3.197a5.971 5.971 0 0 0-.94 3.197M15 6.75a3 3 0 1 1-6 0 3 3 0 0 1 6 0Zm6 3a2.25 2.25 0 1 1-4.5 0 2.25 2.25 0 0 1 4.5 0Zm-13.5 0a2.25 2.25 0 1 1-4.5 0 2.25 2.25 0 0 1 4.5 0Z"/>
                    </svg>

                    <span className="btm-nav-label">Liity kortille</span>
                </Link>

                <Link to={"uusi-peli"} className="">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5"
                         stroke="currentColor" className="w-6 h-6">
                        <path strokeLinecap="round" strokeLinejoin="round" d="M12 4.5v15m7.5-7.5h-15"/>
                    </svg>
                    <span className="btm-nav-label">Uusi tuloskortti</span>
                </Link>
            </nav>
        </>
    )
}

export default App
