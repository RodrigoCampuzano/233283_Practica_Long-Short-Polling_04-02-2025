import React, { useState } from 'react';
import AddProductForm from './components/AddProductForm';
import DiscountCounter from './components/DiscountCounter';
import LastAddedProduct from './components/LastAddedProduct';
import './styles.css';

function App() {
  const [products, setProducts] = useState([]);

  const handleProductAdded = () => {
    // Actualizar la lista de productos cuando se añade uno nuevo
    setProducts(prev => {
      // Simular obtención del último producto (en realidad deberías hacer una petición)
      const newProduct = {
        nombre: 'Nuevo producto',
        precio: '100',
        codigo: 'ABC123',
        descuento: false
      };
      return [newProduct, ...prev].slice(0, 5); // Mantener solo los últimos 5 productos
    });
  };

  return (
    <div className="app-container">
      <h1>Product Management System</h1>
      <div className="components-grid">
        <div className="form-section">
          <AddProductForm onProductAdded={handleProductAdded} />
        </div>
        <div className="count-section">
          <DiscountCounter />
        </div>
        <div className="products-section">
          <LastAddedProduct products={products} />
        </div>
      </div>
    </div>
  );
}

export default App;