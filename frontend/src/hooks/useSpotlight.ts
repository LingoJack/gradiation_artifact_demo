import { useCallback, useRef, useState } from 'react';

interface SpotlightPosition {
  x: number;
  y: number;
}

/**
 * 液态玻璃鼠标跟随光效 Hook
 * 用于实现鼠标移动时的动态光效
 */
export const useSpotlight = () => {
  const [position, setPosition] = useState<SpotlightPosition>({ x: 0, y: 0 });
  const [isHovered, setIsHovered] = useState(false);
  const ref = useRef<HTMLElement>(null);

  const handleMouseMove = useCallback((e: React.MouseEvent<HTMLElement>) => {
    if (!ref.current) return;
    
    const rect = ref.current.getBoundingClientRect();
    setPosition({
      x: e.clientX - rect.left,
      y: e.clientY - rect.top,
    });
  }, []);

  const handleMouseEnter = useCallback(() => {
    setIsHovered(true);
  }, []);

  const handleMouseLeave = useCallback(() => {
    setIsHovered(false);
  }, []);

  const spotlightStyle = {
    '--spotlight-x': `${position.x}px`,
    '--spotlight-y': `${position.y}px`,
    '--spotlight-opacity': isHovered ? 1 : 0,
  } as React.CSSProperties;

  return {
    ref,
    spotlightStyle,
    isHovered,
    handlers: {
      onMouseMove: handleMouseMove,
      onMouseEnter: handleMouseEnter,
      onMouseLeave: handleMouseLeave,
    },
  };
};
