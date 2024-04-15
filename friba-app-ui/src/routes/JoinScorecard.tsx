export default function JoinScorecard() {
    return <div className="h-[80svh] flex flex-col gap-4 mx-4 justify-center">
        <form action="" className="flex flex-col gap-6 justify-center">
            <label className="input input-bordered flex items-center gap-2 input-lg">
                Koodi
                <input type="text" className="grow" placeholder="FR1B4" />
            </label>
            <button className="btn btn-accent self-end">Liity</button>
        </form>
    </div>
}