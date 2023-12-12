import { Button } from "./ui/button"
import { Input } from "./ui/input"
import { Label } from "./ui/label"

export default function JoinScorecard() {
  return (
    <>
      <form className="flex flex-wrap w-full gap-2" action="">
        <Label className="block w-full ">
          <p className="mb-2">Tuloskortin koodi</p>
          <Input />
        </Label>
        <div className="flex w-full justify-end">
          <Button>Liity</Button>
        </div>
      </form>
    </>
  )
}
