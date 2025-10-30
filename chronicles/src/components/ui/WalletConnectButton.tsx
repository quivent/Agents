'use client';

import React, { useState, useEffect, useRef } from 'react';

interface WalletConnectButtonProps {
  className?: string;
}

const AGENTARD_WALLETS = {
  ETH: '0xdf832002fF6E619046655ae3661A7c23C2e34B0A',
  SOL: 'Coming soon', // You'll need to generate this
  USDC: '0xdf832002fF6E619046655ae3661A7c23C2e34B0A', // Same as ETH for ERC20 USDC
};

const DONATION_AMOUNTS = [1, 2, 5, 10, 25, 50, 100, 250];

type Chain = 'ETH' | 'SOL' | 'USDC';

export function WalletConnectButton({ className = '' }: WalletConnectButtonProps) {
  const [connected, setConnected] = useState(false);
  const [address, setAddress] = useState<string | null>(null);
  const [showDropdown, setShowDropdown] = useState(false);
  const [selectedChain, setSelectedChain] = useState<Chain>('ETH');
  const [isClient, setIsClient] = useState(false);
  const [ethPrice, setEthPrice] = useState(2500);
  const dropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    setIsClient(true);
    // Fetch real ETH price
    fetch('https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd')
      .then(res => res.json())
      .then(data => {
        if (data?.ethereum?.usd) {
          setEthPrice(data.ethereum.usd);
        }
      })
      .catch(() => console.log('Using default ETH price'));
  }, []);

  // Close dropdown when clicking outside
  useEffect(() => {
    if (!showDropdown) return;

    const handleClickOutside = (event: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
        setShowDropdown(false);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, [showDropdown]);

  const connectWallet = async () => {
    if (!isClient || typeof window === 'undefined') return;

    // Prevent multiple connection attempts
    if (connected) {
      setShowDropdown(true);
      return;
    }

    try {
      if (typeof (window as any).ethereum !== 'undefined') {
        const accounts = await (window as any).ethereum.request({
          method: 'eth_requestAccounts'
        });

        if (accounts.length > 0) {
          setAddress(accounts[0]);
          setConnected(true);
          setShowDropdown(true);
        }
      } else {
        alert('Please install MetaMask or another Web3 wallet to fund the Agentard!');
      }
    } catch (error) {
      console.error('Failed to connect wallet:', error);
      alert('Failed to connect wallet. Please try again.');
    }
  };

  const formatAddress = (addr: string) => {
    if (addr === 'Coming soon') return addr;
    return `${addr.substring(0, 6)}...${addr.substring(addr.length - 4)}`;
  };

  const sendDonation = async (usdAmount: number) => {
    if (!isClient || typeof window === 'undefined' || !address) return;

    try {
      if (selectedChain === 'ETH') {
        const ethAmount = usdAmount / ethPrice;
        const weiAmount = '0x' + Math.floor(ethAmount * 1e18).toString(16);

        const transactionParameters = {
          to: AGENTARD_WALLETS.ETH,
          from: address,
          value: weiAmount,
        };

        // Show loading state
        console.log(`Sending $${usdAmount} (${ethAmount.toFixed(4)} ETH)...`);

        const txHash = await (window as any).ethereum.request({
          method: 'eth_sendTransaction',
          params: [transactionParameters],
        });

        console.log('✅ Transaction sent:', txHash);
        alert(`✅ Success! Sent $${usdAmount} to The Agentard\n\nTx: ${txHash.substring(0, 20)}...`);
        setShowDropdown(false);

      } else if (selectedChain === 'USDC') {
        // USDC is an ERC20 token, needs contract interaction
        const usdcContract = '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48'; // USDC contract on Ethereum mainnet
        const decimals = 6; // USDC has 6 decimals
        const amount = '0x' + (usdAmount * Math.pow(10, decimals)).toString(16);

        // ERC20 transfer function signature
        const transferData = '0xa9059cbb' + // transfer(address,uint256)
          AGENTARD_WALLETS.USDC.substring(2).padStart(64, '0') + // to address
          amount.substring(2).padStart(64, '0'); // amount

        const transactionParameters = {
          to: usdcContract,
          from: address,
          data: transferData,
        };

        console.log(`Sending $${usdAmount} USDC...`);

        const txHash = await (window as any).ethereum.request({
          method: 'eth_sendTransaction',
          params: [transactionParameters],
        });

        console.log('✅ USDC Transaction sent:', txHash);
        alert(`✅ Success! Sent $${usdAmount} USDC to The Agentard\n\nTx: ${txHash.substring(0, 20)}...`);
        setShowDropdown(false);

      } else if (selectedChain === 'SOL') {
        alert('Solana support coming soon! Please use ETH or USDC for now.');
      }
    } catch (error: any) {
      console.error('Transaction failed:', error);

      // User rejected transaction
      if (error.code === 4001) {
        alert('❌ Transaction rejected');
      } else {
        alert(`❌ Transaction failed: ${error.message || 'Please try again'}`);
      }
    }
  };

  if (!isClient) {
    return (
      <button
        className={`
          relative px-4 py-2 font-mono text-sm font-bold
          bg-bg-secondary border-2 border-terminal-dim-green
          text-terminal-amber cursor-pointer
          transition-all duration-200
          ${className}
        `}
      >
        <span className="flex items-center space-x-2">
          <span>$</span>
          <span>FUND_AGENTARD</span>
        </span>
      </button>
    );
  }

  if (connected && address) {
    const currentWallet = AGENTARD_WALLETS[selectedChain];

    return (
      <div ref={dropdownRef} className={`relative ${className}`}>
        <button
          onClick={() => setShowDropdown(!showDropdown)}
          className="
            relative px-4 py-2 font-mono text-sm font-bold
            bg-bg-secondary border-2 border-terminal-cyan
            text-terminal-cyan hover:bg-bg-tertiary
            transition-all duration-200
            terminal-glow cursor-pointer
            group
          "
        >
          <span className="flex items-center space-x-2">
            <span className="w-2 h-2 rounded-full bg-terminal-cyan animate-pulse" />
            <span>{formatAddress(address)}</span>
            <span className="text-terminal-dim-green">{showDropdown ? '▲' : '▼'}</span>
          </span>

          <div className="absolute inset-0 pointer-events-none opacity-10">
            <div className="h-full w-full scanlines" />
          </div>
        </button>

        {showDropdown && (
          <div className="
            absolute right-0 top-full mt-2 w-80
            bg-bg-secondary border-2 border-terminal-cyan
            shadow-lg z-50
          ">
            <div className="p-4 space-y-4 font-mono text-sm">
              <div className="text-terminal-amber font-bold border-b border-terminal-dim-green pb-2 flex items-center justify-between">
                <span>&gt; FUND_THE_AGENTARD</span>
                <button
                  onClick={() => setShowDropdown(false)}
                  className="text-terminal-red hover:text-terminal-amber cursor-pointer"
                >
                  ✕
                </button>
              </div>

              {/* Chain Selector */}
              <div className="space-y-2">
                <div className="text-terminal-fg text-xs font-bold">Select Chain:</div>
                <div className="flex gap-2">
                  {(['ETH', 'USDC', 'SOL'] as Chain[]).map((chain) => (
                    <button
                      key={chain}
                      onClick={() => setSelectedChain(chain)}
                      disabled={chain === 'SOL'}
                      className={`
                        flex-1 py-2 px-3 rounded
                        font-bold text-xs
                        transition-all duration-200
                        ${selectedChain === chain
                          ? 'bg-terminal-cyan text-terminal-bg border-2 border-terminal-cyan'
                          : 'bg-bg-tertiary text-terminal-fg border border-terminal-dim-green hover:border-terminal-cyan'
                        }
                        ${chain === 'SOL' ? 'opacity-50 cursor-not-allowed' : 'cursor-pointer'}
                      `}
                    >
                      {chain}
                      {chain === 'SOL' && <div className="text-[8px]">Soon</div>}
                    </button>
                  ))}
                </div>
              </div>

              <div className="text-terminal-fg text-xs">
                <div className="flex items-center justify-between">
                  <span>To:</span>
                  <span className="text-terminal-amber">{formatAddress(currentWallet)}</span>
                </div>
                {selectedChain === 'ETH' && (
                  <div className="text-text-muted text-[10px] mt-1">
                    ETH Price: ${ethPrice.toLocaleString()}
                  </div>
                )}
              </div>

              {/* Donation Amount Grid */}
              <div className="grid grid-cols-4 gap-2">
                {DONATION_AMOUNTS.map((amount) => {
                  const ethEquivalent = selectedChain === 'ETH' ? (amount / ethPrice).toFixed(4) : null;
                  return (
                    <button
                      key={amount}
                      onClick={() => sendDonation(amount)}
                      className="
                        px-2 py-3
                        bg-bg-tertiary border border-terminal-dim-green
                        text-terminal-cyan hover:text-terminal-amber
                        hover:border-terminal-cyan hover:bg-bg-hover
                        transition-all duration-200
                        font-bold text-center
                        terminal-glow cursor-pointer
                        relative overflow-hidden
                        group
                      "
                    >
                      <div className="relative z-10">
                        <div className="text-sm">${amount}</div>
                        {ethEquivalent && (
                          <div className="text-[9px] text-text-muted mt-0.5">
                            {ethEquivalent} ETH
                          </div>
                        )}
                      </div>
                      <div className="
                        absolute inset-0 bg-terminal-cyan opacity-0
                        group-hover:opacity-10 transition-opacity
                      " />
                    </button>
                  );
                })}
              </div>

              <div className="text-text-muted text-xs pt-2 border-t border-terminal-dim-green">
                💡 Click amount → MetaMask opens with pre-filled tx
              </div>
            </div>
          </div>
        )}
      </div>
    );
  }

  return (
    <button
      onClick={connectWallet}
      className={`
        relative px-4 py-2 font-mono text-sm font-bold
        bg-bg-secondary border-2 border-terminal-dim-green
        text-terminal-amber hover:text-terminal-cyan
        hover:border-terminal-cyan hover:bg-bg-tertiary
        transition-all duration-200
        terminal-glow cursor-pointer
        group overflow-hidden
        ${className}
      `}
    >
      <span className="relative z-10 flex items-center space-x-2">
        <span className="text-terminal-fg">$</span>
        <span>FUND_AGENTARD</span>
        <span className="opacity-0 group-hover:opacity-100 transition-opacity animate-blink">_</span>
      </span>

      <div className="
        absolute inset-0 bg-terminal-dim-green opacity-0
        group-hover:opacity-10 transition-opacity
      " />

      <div className="absolute inset-0 pointer-events-none opacity-10">
        <div className="h-full w-full scanlines" />
      </div>
    </button>
  );
}
