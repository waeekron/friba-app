import CreateScorecard from "@/components/CreateScorecard"
import JoinScorecard from "@/components/JoinScorecard"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"

export default function Root() {
  return (
    <>
      <div className="flex justify-between mx-4">
        <h1 className="font-bold text-xl">Tuloskortti</h1> <div>settings</div>
      </div>
      <div className="h-full flex items-center justify-center">
        <div className="flex flex-wrap justify-center gap-4 my-10 max-w-sm ">
          <Tabs defaultValue="create" className="w-[400px]">
            <TabsList className="grid w-full grid-cols-2">
              <TabsTrigger value="create">Luo uusi tuloskortti</TabsTrigger>
              <TabsTrigger value="join">Liity tuloskortille</TabsTrigger>
            </TabsList>
            <TabsContent value="create">
              <div className="h-96 w-full px-2">
                <CreateScorecard />
              </div>
            </TabsContent>
            <TabsContent value="join">
              <div className="h-96 w-full px-2">
                <JoinScorecard></JoinScorecard>
              </div>
            </TabsContent>
          </Tabs>
        </div>
      </div>
    </>
  )
}
