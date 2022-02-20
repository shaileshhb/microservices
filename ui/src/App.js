import CreatePost from "./CreatePost";
import DisplayPost from "./DisplayPost";

function App() {
  return(
    <div className="container">
      <h1>Blog App</h1>
      <CreatePost/>
      <DisplayPost/>
    </div>
  )
}

export default App;