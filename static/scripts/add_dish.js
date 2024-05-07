document.addEventListener('DOMContentLoaded', function() {
    document.querySelector('form').addEventListener('submit', function(e) {
        if (!validateForm()) {
            e.preventDefault(); // Предотвращаем отправку формы, если валидация не прошла
        }
    });
});

function validateForm() {
    var name = document.getElementById('dish_name').value;
    var category = document.getElementById('dish_category').value;
    var description = document.getElementById('dish_description').value;
    var price = document.getElementById('dish_price').value;
    var isValid = true;

    // Проверка поля "Name"
    if (name === '') {
        document.getElementById('name-error').innerHTML = 'Please enter a name.';
        document.getElementById('dish_name').value = '';
        isValid = false;
    } else {
        document.getElementById('name-error').innerHTML = '';
    }

    // Проверка поля "Category"
    if (category === '') {
        document.getElementById('category-error').innerHTML = 'Please select a category.';
        document.getElementById('dish_category').value = '';
        isValid = false;
    } else {
        document.getElementById('category-error').innerHTML = '';
    }

    // Проверка поля "Description"
    if (description === '') {
        document.getElementById('description-error').innerHTML = 'Please enter a description.';
        document.getElementById('dish_description').value = '';
        isValid = false;
    } else {
        document.getElementById('description-error').innerHTML = '';
    }

    // Проверка поля "Price"
    if (isNaN(price) || price === '') {
        document.getElementById('price-error').innerHTML = 'Please enter a valid price.';
        document.getElementById('dish_price').value='';
        isValid = false;
    } else {
        document.getElementById('price-error').innerHTML = '';
    }

    return isValid; // Возвращаем true, если все поля прошли валидацию, иначе false
}
