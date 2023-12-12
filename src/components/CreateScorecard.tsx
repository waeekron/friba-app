import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { ScrollArea } from "@/components/ui/scroll-area"
import { Separator } from "@/components/ui/separator"
import { useState } from "react"
export default function CreateScorecard() {
  return (
    <>
      <CourseList />
    </>
  )
}

const courseList = Array.from({ length: 50 }).map(
  (_, i, a) => `firsbeegolf kenttä ${a.length - i}`
)
courseList.push("Laahalahti")
function CourseList() {
  const [courseSearch, setCourseSearch] = useState("")
  const [selectedCourse, setSelectedCourse] = useState("")
  const [username, setUsername] = useState("")
  const [courses, setCourses] = useState(courseList)
  const [showPreview, setShowPreview] = useState(false)

  return (
    <>
      {!showPreview ? (
        <div className="flex flex-wrap">
          <div className="h-96 w-full flex gap-1 flex-wrap">
            <Input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              placeholder="Käyttäjänimi"
            />
            <Input
              type="text"
              value={courseSearch}
              onChange={(e) => {
                setCourseSearch(e.target.value)
                console.log(e.target.value)
                // check if searched course exists
                setSelectedCourse(e.target.value)
                const filtered = courseList.filter((c) =>
                  c.toLowerCase().includes(e.target.value.toLowerCase())
                )
                setCourses(filtered)
              }}
              placeholder="Hae rataa"
            />
            <ScrollArea className="h-64 w-full   rounded-md border">
              <div className="p-4">
                <h4 className="mb-4 text-sm font-medium leading-none">
                  Valitse rata
                </h4>
                {courses.map((c) => (
                  <>
                    <div
                      key={c}
                      className={`text-sm py-2 ${
                        selectedCourse === c ? "bg-blue-100" : "none"
                      }`}
                      onClick={() => {
                        console.log()
                        setSelectedCourse(c)
                        setCourseSearch(c)
                      }}
                      role="button"
                    >
                      {c}
                    </div>
                    <Separator className="my-1" />
                  </>
                ))}
              </div>
            </ScrollArea>
          </div>
        </div>
      ) : (
        <div className="flex gap-2 flex-wrap">
          <div className="h-96"></div>
        </div>
      )}
      <div
        className={`flex justify-end ${
          showPreview && "justify-between"
        } w-full mt-4`}
      >
        {showPreview && (
          <Button
            className="w-24"
            disabled={!selectedCourse || !username}
            onClick={() => setShowPreview(!showPreview)}
          >
            Edellinen
          </Button>
        )}
        <Button
          className="w-24"
          disabled={!selectedCourse || !username}
          onClick={
            showPreview
              ? () => alert("Aloitetaan")
              : () => setShowPreview(!showPreview)
          }
        >
          {showPreview ? "Aloita" : "Jatka"}
        </Button>
      </div>
    </>
  )
}
