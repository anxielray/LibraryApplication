// Toggle dropdown menu visibility
function toggleMenu() {
    const menu = document.getElementById('dropdownMenu');
    menu.style.display = menu.style.display === 'block' ? 'none' : 'block';
  }
  
  // Handle action when a menu item is clicked
  function handleAction(action) {
    alert(`${action} option selected!`);
    // Add logic for each action here, like navigation or API calls.
  }
  