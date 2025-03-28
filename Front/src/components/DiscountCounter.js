import React, { useState, useEffect } from 'react';
import axios from 'axios';

const DiscountCounter = () => {
  const [count, setCount] = useState(0);
  const [isLoading, setIsLoading] = useState(true);

  const fetchDiscountCount = async () => {
    try {
      const response = await axios.get('http://localhost:8080/CountProductIsInDiscount');
      setCount(response.data.count);
    } catch (error) {
      console.error('Error fetching discount count:', error);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchDiscountCount();
    const interval = setInterval(fetchDiscountCount, 3000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="discount-counter">
      <h2>Products on Discount</h2>
      {isLoading ? (
        <p>Loading count...</p>
      ) : (
        <div className="count-display">
          <span className="count-number">{count}</span>
          <span className="count-label">items</span>
        </div>
      )}
    </div>
  );
};

export default DiscountCounter;