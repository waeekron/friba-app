import {Link} from "react-router-dom";

const courses = [{
    name: "Laajalahti",
    city: "Kokkola"
}]
export default function CreateScorecard() {
    function handleRowClick(e) {


    }

    return <div className="h-[80svh] flex flex-col gap-4 mx-4 items-center">

        <form action="">
            <label className="input input-bordered flex items-center gap-2">
                <input type="text" className="grow input-lg" placeholder="Search"/>
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor"
                     className="w-4 h-4 opacity-70">
                    <path fillRule="evenodd"
                          d="M9.965 11.026a5 5 0 1 1 1.06-1.06l2.755 2.754a.75.75 0 1 1-1.06 1.06l-2.755-2.754ZM10.5 7a3.5 3.5 0 1 1-7 0 3.5 3.5 0 0 1 7 0Z"
                          clipRule="evenodd"/>
                </svg>
            </label>
        </form>
        <div className="overflow-x-auto h-[60svh] w-full">
            <table className="table table-pin-rows  table-lg">
                <thead className="">
                <tr>
                    <th>Nimi</th>
                    <th>Kaupunki</th>
                </tr>
                </thead>
                <tbody>
                {Array(40).fill({nimi: "laajis", kaupunki: "kokkola"}).map(item =>
                    (<tr onClick={(e) => handleRowClick(e)}>
                        <td>{item.nimi}</td>
                        <td>{item.kaupunki}</td>
                    </tr>))}
                <tr>
                    <td>Laajalahti</td>
                    <td>Kokkola</td>
                </tr>
                </tbody>
            </table>
        </div>

        <Link className="flex justify-end w-full" to={"/tuloskortti/:koodi"}>
            <button className="btn btn-accent  rounded-bl-btn">Luo</button>
        </Link>
    </div>
}