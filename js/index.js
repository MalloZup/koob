function sendDataToBackend(){
   // Get the data from each element on the form.
   // It is a bit hardcoded  
   const name = document.getElementById('linkName');
   const tag = document.getElementById('linkTag');
   // not satisfied by this name, maybe improve later
   linkInfos = { "name": name.value, "tag": tag.value };
   console.log(linkInfos);
}
