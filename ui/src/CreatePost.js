import { useState } from "react";
import axios from "axios";

function CreatePost() {
  const titleLabel = "Enter Post Title";
  const [postTitle, updatePostTitle] = useState("");

  const onSubmitHandler = async (e) => {
    e.preventDefault();
    await axios
      .post("http://localhost:4001/api/v1/posts", {
        title: postTitle,
      })
      .catch((err) => console.error(err));

    updatePostTitle("");
  };

  const onSubmitClick = () => {
    if (postTitle === "" || postTitle === null) {
      alert("Post title must be specified")
      return;
    }
    updatePostTitle(postTitle);
  };

  return (
    <div className="container-fluid">
      {/* onSubmit={onSubmitHandler} --> doesn't work for me */}
      <form onSubmit={onSubmitHandler}>
        <div className="row">
          <div className="col-sm-6 form-group">
            <label>{titleLabel}:</label>
            <input type="text" className="form-control" placeholder="Post-1" 
             value={postTitle} onChange={(e) => updatePostTitle(e.target.value)}/>
          </div>
        </div>
        <br />
        <button type="submit" className="btn btn-primary" onClick={onSubmitClick}>
          Submit
        </button>
      </form>
      <br />
    </div>
  );
}

export default CreatePost;
