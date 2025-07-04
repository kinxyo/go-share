const folderSelect = document.getElementById('folder-select');
const fileList = document.getElementById('file-list');

function clearFiles() {
  fileList.innerText = 'Select a folder...';
}

function loadFolders() {
  fetch('/api/folders')
    .then(res => res.json())
    .then(folders => {
      folderSelect.innerHTML = '<option value="" disabled selected>Select folder</option>';
      folders.forEach(folder => {
        const opt = document.createElement('option');
        opt.value = folder;
        opt.textContent = folder.charAt(0).toUpperCase() + folder.slice(1);
        folderSelect.appendChild(opt);
      });
    })
    .catch(() => {
      folderSelect.innerHTML = '<option value="" disabled>Error loading folders</option>';
    });
}

folderSelect.addEventListener('change', () => {
  const folder = folderSelect.value;
  if (!folder) {
    clearFiles();
    return;
  }
  fetch(`/api/files?folder=${folder}`)
    .then(res => res.json())
    .then(files => {
      fileList.innerHTML = '';
      if (files.length === 0) {
        fileList.innerText = 'No files found in this folder.';
        return;
      }
      files.forEach(file => {
        const link = document.createElement('a');
        link.href = `/files/${folder}/${encodeURIComponent(file)}`;
        link.innerText = file;
        link.download = file;
        fileList.appendChild(link);
      });
    })
    .catch(() => {
      fileList.innerText = 'Error loading files';
    });
});

// Load folders on page load
loadFolders();
clearFiles();

