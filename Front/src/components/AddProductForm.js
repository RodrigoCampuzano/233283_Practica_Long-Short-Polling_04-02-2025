import React, { useState } from 'react';
import axios from 'axios';

const AddProductForm = () => {
  const [product, setProduct] = useState({
    name: '',
    price: '',
    code: '',
    onDiscount: false
  });
  const [isLoading, setIsLoading] = useState(false);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setProduct(prev => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    if (!product.name || !product.price || !product.code) {
      alert('Please fill all fields');
      return;
    }
    
    setIsLoading(true);
    try {
      await axios.post('http://localhost:8080/addProduct', {
        nombre: product.name,
        precio: product.price,
        codigo: product.code,
        descuento: product.onDiscount
      });
      setProduct({
        name: '',
        price: '',
        code: '',
        onDiscount: false
      });
      alert('Product added successfully!');
    } catch (error) {
      console.error('Error adding product:', error);
      alert('Error adding product. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="product-form">
      <h2>Add New Product</h2>
      <div className="form-group">
        <label>Name:</label>
        <input
          type="text"
          name="name"
          value={product.name}
          onChange={handleChange}
          required
        />
      </div>
      <div className="form-group">
        <label>Price ($):</label>
        <input
          type="number"
          name="price"
          value={product.price}
          onChange={handleChange}
          required
          min="0"
          step="0.01"
        />
      </div>
      <div className="form-group">
        <label>Product Code:</label>
        <input
          type="text"
          name="code"
          value={product.code}
          onChange={handleChange}
          required
        />
      </div>
      <div className="form-group checkbox">
        <label>
          <input
            type="checkbox"
            name="onDiscount"
            checked={product.onDiscount}
            onChange={handleChange}
          />
          On Discount
        </label>
      </div>
      <button type="submit" disabled={isLoading}>
        {isLoading ? 'Adding...' : 'Add Product'}
      </button>
    </form>
  );
};

export default AddProductForm;