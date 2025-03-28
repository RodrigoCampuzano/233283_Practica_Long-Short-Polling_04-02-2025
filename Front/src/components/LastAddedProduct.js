import React, { useState, useEffect } from 'react';
import axios from 'axios';

const LastAddedProduct = () => {
  const [products, setProducts] = useState([]);
  const [isLoading, setIsLoading] = useState(true);

  const fetchLastProduct = async () => {
    try {
      const response = await axios.get('http://localhost:8080/isNewProductAdded');
      if (response.data.hasNewProduct && response.data.product) {
        setProducts(prev => [response.data.product, ...prev].slice(0, 5)); // Mantener solo 5 últimos
      }
    } catch (error) {
      console.error('Error fetching last product:', error);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    // Primera carga
    fetchLastProduct();
    
    // Configurar polling cada 1 segundo
    const intervalId = setInterval(fetchLastProduct, 1000);
    
    return () => clearInterval(intervalId);
  }, []);

  if (isLoading) {
    return <p>Loading...</p>;
  }

  return (
    <div className="last-product">
      <h2>Last Added Products</h2>
      {products.length === 0 ? (
        <p>No products added yet</p>
      ) : (
        <div className="products-list">
          {products.map((product, index) => (
            <div key={product.codigo + index} className="product-details">
              <h3>Product #{index + 1}</h3>
              <p><strong>Name:</strong> {product.nombre}</p>
              <p><strong>Price:</strong> {product.precio}</p>
              <p><strong>Code:</strong> {product.codigo}</p>
              <p><strong>On Discount:</strong> {product.descuento ? 'Yes' : 'No'}</p>
              <hr />
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default LastAddedProduct;