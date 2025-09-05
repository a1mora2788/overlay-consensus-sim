import pandas as pd
import matplotlib.pyplot as plt
import argparse

# --- Arguments ---
parser = argparse.ArgumentParser(description="Plot Overlay vs PoW Energy Ratios (pitch-ready)")
parser.add_argument("--windows", type=int, nargs="+", default=[10],
                    help="List of smoothing window sizes (default: 10)")
parser.add_argument("--raw", action="store_true",
                    help="Include raw per-epoch ratio (noisy)")
args = parser.parse_args()

# --- Load CSV ---
df = pd.read_csv("per_epoch_results.csv")
df.columns = df.columns.str.strip()

# --- Compute ratios ---
df["overlay_vs_pow_ratio"] = df["overlay_energy"] / df["pow_energy"].replace(0, float("nan"))
df["cum_overlay_vs_pow_ratio"] = df["cum_overlay_energy"] / df["cum_pow_energy"].replace(0, float("nan"))

# --- Rolling averages ---
for w in args.windows:
    df[f"overlay_vs_pow_ratio_smooth_{w}"] = df["overlay_vs_pow_ratio"].rolling(
        window=w, min_periods=1, center=True
    ).mean()

# --- Save updated CSV ---
df.to_csv("per_epoch_results_with_ratios.csv", index=False)

# --- Plot ---
plt.figure(figsize=(12, 7))

# Optional raw line
if args.raw:
    plt.plot(df["epoch"], df["overlay_vs_pow_ratio"],
             label="Overlay/PoW Energy Ratio (per epoch)", color="blue", alpha=0.3)

# Smoothed line(s)
for w in args.windows:
    plt.plot(df["epoch"], df[f"overlay_vs_pow_ratio_smooth_{w}"],
             label=f"Smoothed Ratio ({w} epochs)", linewidth=2)

# Cumulative ratio
plt.plot(df["epoch"], df["cum_overlay_vs_pow_ratio"],
         label="Cumulative Overlay/PoW Energy Ratio", color="green", linewidth=3)

plt.xlabel("Epoch", fontsize=12)
plt.ylabel("Energy Ratio", fontsize=12)
plt.title("Overlay vs PoW Energy Ratios", fontsize=14, weight="bold")
plt.legend(loc="lower right", fontsize=10, frameon=False)
plt.grid(True, alpha=0.3)
plt.tight_layout()

# Save figures
plt.savefig("overlay_vs_pow.png", dpi=300)
plt.savefig("overlay_vs_pow.pdf")  # good for presentations/papers
plt.show()

print("Pitch-ready plots saved as overlay_vs_pow.png and overlay_vs_pow.pdf")
