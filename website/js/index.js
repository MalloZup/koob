function sendDataToBackend(){
   // Get the data from each element on the form.
   // It is a bit hardcoded  
   const Name = document.getElementById('linkName');
   const Tag = document.getElementById('linkTag');
   // not satisfied by this name, maybe improve later
   let linkInfos = { "Name": Name.value, "Tag": Tag.value };
   console.log(linkInfos);
   console.log("LOGGED");
   fetch("/bookmarks", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },

      body: JSON.stringify(linkInfos),
    }).then((resp) => {
      console.log(resp);
    });
}
