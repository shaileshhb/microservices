import { useEffect, useState } from "react";
import axios from "axios";


function DisplayPost() {

  const [posts, updatePosts] = useState({})

  const loadPosts = async () => {
    const response = await axios.get("http://localhost:4001/api/v1/posts").
      catch(err => console.error(err))

      console.log(response);
      updatePosts(response.data)
  }

  useEffect(() => {
    loadPosts()
  }, []) // [] -> callback only once

  const cardOfPosts = Object.values(posts).map(p => {
    return (
      <div className="card" key={p.id}>
        <div className="card-body">
          {p.title}
        </div>
      </div>
    )
  })

  return (
    <div className="container">
      {/* <div className="card"></div> */}
      {cardOfPosts}
    </div>
  )
}

export default DisplayPost;