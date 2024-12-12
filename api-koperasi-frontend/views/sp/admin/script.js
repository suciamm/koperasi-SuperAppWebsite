
//ACTION DROPDOWN ICON PERSON
document.addEventListener("DOMContentLoaded", function () {
    const profileIcon = document.querySelector(".profile-icon");
    const profileDropdown = document.querySelector(".profile-dropdown");

    profileIcon.addEventListener("click", function () {
        profileDropdown.classList.toggle("active");
    });

    // Menutup dropdown jika klik di luar
    document.addEventListener("click", function (event) {
        if (!profileIcon.contains(event.target)) {
            profileDropdown.classList.remove("active");
        }
    });
});


//ICON MENU PROFILE
    // fetch('http://localhost:8080/api/dataIconPerson')
    //     .then(response => response.json())
    //     .then(data =>{
    //         const tableBody = document.getElementById('profile-dropdown');
    //         data.forEach(profile => {
    //             const row = document.createElement('tr');
    //             row.innerHTML = `
    //                 <h5>${profile.email}</h5>
    //                 <p>${profile.verified}</p>  
    //         `;
    //         tableBody.appendChild(row);
    //         });
    //     })
    // .catch(error => {
    //     console.error('Error fetching data:', error);
    //     alert('Gagal ambil datanyaaaa');
    // });