import { useState } from "react";

function CreatePost() {

  const titleLabel = "Enter Post Title"
  const [postTitle, updatePostTitle] = useState("")

  // const onSubmitHandler = (e) => {
  //   console.log(e);
  // }

  const onSubmitClick = () => {
    if (postTitle == "" || postTitle == null) {
      // alert("Post title must be specified")
      return
    }
    console.log(postTitle);
    updatePostTitle("")
  }

  return (
    <div className="container-fluid">
      {/* onSubmit={onSubmitHandler} --> doesn't work for me */}
      <form>
        <div className="row">
          <div className="col-sm-6 form-group">
            <label>{titleLabel}:</label>
            <input type="text" className="form-control" placeholder="Post-1"
              value={postTitle} onChange={(e)=>updatePostTitle(e.target.value)}/>
          </div>
        </div>
      </form>
      <br/>
      <button type="submit" className="btn btn-primary" onClick={onSubmitClick}>Submit</button>
      <br/>
      {postTitle}
    </div>
  );
}

export default CreatePost;
