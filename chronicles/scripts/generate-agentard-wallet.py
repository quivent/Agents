#!/usr/bin/env python3
"""
Generate an Ethereum wallet for The Agentard using Wallefestor
"""
import sys
import os
import json
from datetime import datetime
from pathlib import Path

# Add Wallefestor to path
sys.path.insert(0, str(Path.home() / 'Wallefestor'))

from wallefestor.core.factory import WalletFactory
from wallefestor.chains.ethereum import EthereumImplementation
import secrets

def generate_agentard_wallet():
    """Generate a secure Ethereum wallet for The Agentard"""

    print("🤖 Generating Ethereum wallet for THE AGENTARD...")
    print("=" * 60)

    # Generate secure entropy
    entropy = secrets.token_bytes(32)  # 256 bits of secure random data

    # Initialize wallet factory
    factory = WalletFactory()

    # Generate wallet with custom entropy (testnet=False means mainnet)
    wallet = factory.create_wallet(
        chain='ethereum',
        testnet=False,
        entropy=entropy
    )

    # Extract wallet details
    wallet_info = {
        'name': 'THE AGENTARD',
        'purpose': 'Fund the Agentard - Supporting confident incompetence',
        'chain': 'Ethereum',
        'network': 'mainnet',
        'address': wallet.address,
        'label': 'THE_AGENTARD_FUND',
        'derivation_path': "m/44'/60'/0'/0/0",
        'created_at': datetime.now().isoformat()
    }

    # Sensitive data (private key) - to be stored securely
    # Convert bytes to hex string for JSON serialization
    sensitive_data = {
        'private_key': wallet.private_key.hex() if isinstance(wallet.private_key, bytes) else wallet.private_key,
        'public_key': wallet.public_key.hex() if hasattr(wallet, 'public_key') and isinstance(wallet.public_key, bytes) else (wallet.public_key if hasattr(wallet, 'public_key') else '')
    }

    # Print public wallet info
    print("\n✅ WALLET GENERATED SUCCESSFULLY")
    print("=" * 60)
    print(f"💼 Wallet Name: {wallet_info['name']}")
    print(f"🎯 Purpose: {wallet_info['purpose']}")
    print(f"⛓️  Chain: {wallet_info['chain']} ({wallet_info['network']})")
    print(f"📫 Address: {wallet_info['address']}")
    print(f"🏷️  Label: {wallet_info['label']}")
    print(f"🛤️  Derivation Path: {wallet_info['derivation_path']}")
    print("=" * 60)

    # Save public wallet info (safe to commit)
    output_dir = Path(__file__).parent.parent / 'public'
    output_dir.mkdir(exist_ok=True)

    public_wallet_file = output_dir / 'agentard-wallet.json'
    with open(public_wallet_file, 'w') as f:
        json.dump(wallet_info, f, indent=2)

    print(f"\n📄 Public wallet info saved to: {public_wallet_file}")

    # Save sensitive data securely (should NOT be committed)
    secure_dir = Path.home() / '.agentard-wallet'
    secure_dir.mkdir(mode=0o700, exist_ok=True)

    secure_file = secure_dir / 'agentard-wallet-secure.json'
    with open(secure_file, 'w') as f:
        json.dump({**wallet_info, **sensitive_data}, f, indent=2)

    os.chmod(secure_file, 0o600)

    print(f"🔐 Secure wallet data saved to: {secure_file}")
    print(f"⚠️  IMPORTANT: Keep {secure_file} private and secure!")
    print("\n" + "=" * 60)
    print("🚀 Wallet ready to receive funds!")
    print("💰 Send ETH or ERC20 tokens to:", wallet_info['address'])
    print("=" * 60)

    return wallet_info

if __name__ == '__main__':
    try:
        wallet_info = generate_agentard_wallet()

        print("\n🎉 Success! The Agentard now has an Ethereum wallet.")
        print("💡 Next steps:")
        print("   1. Verify the address on Etherscan")
        print("   2. Test with a small amount first")
        print("   3. Update the website wallet connect button")
        print("   4. BACKUP the mnemonic phrase securely")

    except Exception as e:
        print(f"\n❌ Error generating wallet: {e}")
        import traceback
        traceback.print_exc()
        sys.exit(1)
