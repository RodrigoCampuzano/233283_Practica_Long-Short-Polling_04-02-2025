import React, { useState, useEffect, useCallback } from 'react';
import axios from 'axios';

const DiscountCounter = () => {
    const [count, setCount] = useState(0);
    const [isLoading, setIsLoading] = useState(true);

    const fetchDiscountCount = useCallback(async () => {
        try {
            const response = await axios.get('http://localhost:8080/CountProductIsInDiscount', {
                timeout: 5000
            });
            setCount(response.data.count);
        } catch (error) {
            console.error('Error fetching discount count:', error);
        } finally {
            setIsLoading(false);
        }
    }, []);

    useEffect(() => {
        let isMounted = true;
        let timeoutId;

        const poll = async () => {
            if (!isMounted) return;
            await fetchDiscountCount();
            timeoutId = setTimeout(poll, 3000); // Reducido a 3 segundos
        };

        poll();

        return () => {
            isMounted = false;
            clearTimeout(timeoutId);
        };
    }, [fetchDiscountCount]);

    return (
        <form onSubmit={handleSubmit} className="product-form">
          <h2>Add New Product</h2>
          <div className="form-group">
            <label>Name:</label>
            <input
              type="text"
              name="nombre"
              value={product.nombre}
              onChange={handleChange}
              required
            />
          </div>
          <div className="form-group">
            <label>Price:</label>
            <input
              type="number"
              name="precio"
              value={product.precio}
              onChange={handleChange}
              required
              min="0"
              step="0.01"
            />
          </div>
          <div className="form-group">
            <label>Code:</label>
            <input
              type="text"
              name="codigo"
              value={product.codigo}
              onChange={handleChange}
              required
            />
          </div>
          <div className="form-group checkbox">
            <label>
              <input
                type="checkbox"
                name="descuento"
                checked={product.descuento}
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

export default React.memo(DiscountCounter);



