#!/usr/bin/env bash
# Usage: ./trivy-image.sh <image[:tag]> [output_dir]
set -uo pipefail

IMG="${1:-}"
[ -z "$IMG" ] && { echo "Usage: $0 <image[:tag]> [output_dir]"; exit 1; }

command -v trivy >/dev/null || { echo "trivy not installed"; exit 1; }
command -v jq    >/dev/null || { echo "jq not installed (sudo apt install jq / sudo dnf install jq)"; exit 1; }

SAFE=${IMG//[\/:]/_}
OUT="${2:-$HOME/trivy-reports/${SAFE}_$(date +%F_%H%M)}"
mkdir -p "$OUT"

echo "Scanning $IMG ..."
trivy image --quiet --scanners vuln --format json \
  --output "$OUT/report.json" "$IMG" || { echo "Scan failed"; exit 1; }

# Full readable table (all severities)
trivy image --quiet --scanners vuln --format table \
  --output "$OUT/report.txt" "$IMG"

# CSV of every finding
jq -r '
  ["target","cve","severity","package","installed","fixed_in","title"],
  (.Results[]? | .Target as $t | (.Vulnerabilities // [])[] |
   [$t, .VulnerabilityID, .Severity, .PkgName, .InstalledVersion,
    (.FixedVersion // ""), (.Title // "")]) | @csv
' "$OUT/report.json" > "$OUT/issues.csv"

# Severity counts
echo
echo "=== Severity counts for $IMG ==="
jq -r '
  [.Results[]?.Vulnerabilities // [] | .[]] as $v |
  ["CRITICAL","HIGH","MEDIUM","LOW","UNKNOWN"][] as $s |
  "\($s): \([$v[] | select(.Severity==$s)] | length)"
' "$OUT/report.json"

# Fix-first list
echo
echo "=== CRITICAL/HIGH with a fix available ==="
jq -r '
  [.Results[]?.Vulnerabilities // [] | .[]
   | select((.Severity=="CRITICAL" or .Severity=="HIGH") and (.FixedVersion // "") != "")]
  | sort_by(.Severity) | .[]
  | "\(.Severity)\t\(.PkgName)\t\(.InstalledVersion) -> \(.FixedVersion)\t\(.VulnerabilityID)"
' "$OUT/report.json" | column -t -s $'\t'

echo
echo "Saved to: $OUT"
echo "  report.txt  (readable table)"
echo "  issues.csv  (all findings)"
echo "  report.json (raw data)"
