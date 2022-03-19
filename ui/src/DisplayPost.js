import { useEffect, useState } from "react";
import axios from "axios";
import CreateComment from "./CreateComment";
import DisplayComment from './DisplayComment';

function DisplayPost() {
  const [posts, updatePosts] = useState({});

  const loadPosts = async () => {
    const response = await axios
      .get("http://localhost:4003/api/v1/posts")
      .catch((err) => console.error(err));

    console.log(response);
    updatePosts(response.data);
  };

  useEffect(() => {
    loadPosts();
  }, []); // [] -> callback only once

  const cardOfPosts = Object.values(posts).map((p) => {
    return (
      <div className="d-flex justify-content-between" key={p.id}>
        <div className="card">
          <div className="card-body">
            <h3>{p.title}</h3>
            <DisplayComment comments={p.comments}/>
            <CreateComment postID={p.id}/>
          </div>
        </div>
      </div>
    );
  });

  return <div className="container">{cardOfPosts}</div>;
}

export default DisplayPost;
