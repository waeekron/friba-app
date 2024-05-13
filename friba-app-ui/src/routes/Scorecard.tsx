
function TableView() {
    return (
        <div className=" overflow-x-auto">
            <table className="table table-lg table-pin-rows table-pin-cols">
                <thead>
                <tr>
                    <th>väylä</th>
                    <td>2</td>
                    <td>3</td>
                    <td>4</td>
                    <td>5</td>
                    <td>6</td>
                    <td>7</td>
                    <td>8</td>
                    <td>9</td>
                    <td>10</td>
                    <td>11</td>
                    <td>12</td>
                    <td>13</td>
                    <td>14</td>
                    <td>15</td>
                    <th>16</th>
                    <th>17</th>
                    <th></th>
                </tr>
                <tr>
                    <th>par</th>
                    <td>2</td>
                    <td>3</td>
                    <td>4</td>
                    <td>5</td>
                    <td>3</td>
                    <td>4</td>
                    <td>3</td>
                    <td>3</td>
                    <td>4</td>
                    <td>1</td>
                    <td>2</td>
                    <td>3</td>
                    <td>4</td>
                    <td>5</td>
                    <th>6</th>
                    <th>3</th>
                    <th></th>
                </tr>
                </thead>
                <tbody>
                <tr className="">
                    <th>anni</th>
                    <td>2</td>
                    <td>3</td>
                    <td>4</td>
                    <td>5</td>
                    <td>6</td>
                    <td>7</td>
                    <td>8</td>
                    <td>9</td>
                    <td>10</td>
                    <td>11</td>
                    <td>12</td>
                    <td>13</td>
                    <td>14</td>
                    <td>15</td>
                    <th>16</th>
                    <th>17</th>
                    <th>18</th>
                </tr>
                <tr>
                    <th>sani</th>
                    <td>2</td>
                    <td>3</td>
                    <td>4</td>
                    <td>5</td>
                    <td>6</td>
                    <td>7</td>
                    <td>8</td>
                    <td>9</td>
                    <td>10</td>
                    <td>11</td>
                    <td>12</td>
                    <td>13</td>
                    <td>14</td>
                    <td>15</td>
                    <th>16</th>
                    <th>17</th>
                    <th>18</th>
                </tr>
                <tr>
                    <th>walter</th>
                    <td>2</td>
                    <td>3</td>
                    <td>4</td>
                    <td>5</td>
                    <td>6</td>
                    <td>7</td>
                    <td>8</td>
                    <td>9</td>
                    <td>10</td>
                    <td>11</td>
                    <td>12</td>
                    <td>13</td>
                    <td>14</td>
                    <td>15</td>
                    <th>16</th>
                    <th>17</th>
                    <th>18</th>
                </tr>
                </tbody>

            </table>
        </div>
    )
}

function FairwayView() {
    return (
        <div className="flex flex-col items-center gap-8">
            <h1 className="text-4xl rounded-full bg-accent p-2 px-5 border-black border-4    font-extrabold font-mono">1</h1>
            <ul className="flex flex-col gap-3">
                <FairwayViewitem/>
                <FairwayViewitem/>
                <FairwayViewitem/>
            </ul>
        </div>
    )
}

function FairwayViewitem() {
    return (
        <>
            <div className="flex items-center gap-4  border-rose-500 border-4 rounded-badge">

                <button className="btn rounded-l-full btn-lg font-bold text-3xl ">-</button>
                <span className="text-3xl font-bold ">walter</span>
                <span className="stat-value">3</span>
                <button className="btn rounded-r-full font-bold btn-lg text-3xl">+</button>
            </div>
        </>
    )
}

export default function Scorecard() {
    return (
        <>
            <div className="">
                <div className="overflow-y-scroll">

                    <TableView/>

                    <FairwayView/>
                </div>
                <div className=" bg-base-100 fixed bottom-0  w-full flex join join-horizontal">
                    <button className="btn btn-outline join-item flex flex-grow"> left</button>
                    <button className="btn join-item btn btn-outline flex-grow">Button</button>
                    <button className="btn join-item flex-grow btn btn-outline">Button</button>
                </div>

            </div>

        </>
    )
}