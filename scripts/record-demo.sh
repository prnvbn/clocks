#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
TMP_DIR="$(mktemp -d "${TMPDIR:-/tmp}/clocks-demo.XXXXXX")"
trap 'rm -rf "$TMP_DIR"' EXIT
BIN="$TMP_DIR/clocks"
GO_CACHE="${GOCACHE:-$TMP_DIR/go-build-cache}"
CAST_DIR="${CLOCKS_DEMO_CAST_DIR:-$TMP_DIR}"

for cmd in go expect asciinema agg magick; do
    if ! command -v "$cmd" >/dev/null 2>&1; then
        echo "error: '$cmd' is required to record the demos" >&2
        exit 1
    fi
done

mkdir -p "$ROOT_DIR/demo" "$CAST_DIR"
echo "Building clocks..."
(
    cd "$ROOT_DIR"
    GOCACHE="$GO_CACHE" go build -o "$BIN" .
)

# Use the app's generated defaults without touching the user's config.
CLOCKS_CONFIG_PATH="$TMP_DIR/defaults.yaml" "$BIN" config >/dev/null
SETTINGS="$(sed -n '/^live:/,$p' "$TMP_DIR/defaults.yaml")"

if [[ $# -eq 0 ]]; then
    set -- overview clocks layout-grid layout-centered seconds twenty-four-hour
fi

for name in "$@"; do
    ROWS=22
    COLS=112
    CENTER=true
    ROW_SIZES='[2, 2]'
    ARGS=""
    case "$name" in
        overview) ROW_SIZES='[2, 1]' ;;
        clocks) ;;
        layout-grid) CENTER=false; ROW_SIZES='[2, 1]' ;;
        layout-centered) ROW_SIZES='[2, 1]' ;;
        seconds) ROW_SIZES='[2]'; ROWS=13; ARGS='--seconds' ;;
        twenty-four-hour) ROW_SIZES='[2]'; ROWS=13 ;;
        *) echo "error: unknown demo: $name" >&2; exit 1 ;;
    esac

    CONFIG="$TMP_DIR/$name.yaml"
    cat >"$CONFIG" <<EOF
$SETTINGS
layout:
    layoutType: 2
    rowSizes: $ROW_SIZES
    centerEachRow: $CENTER
clocks:
    - heading: London, UK
      zone: Europe/London
      color: cyan
EOF
    if [[ "$name" != twenty-four-hour ]]; then
        cat >>"$CONFIG" <<'EOF'
    - heading: Abu Dhabi, UAE
      zone: Asia/Dubai
      color: yellow
EOF
    fi
    if [[ "$name" != seconds && "$name" != twenty-four-hour ]]; then
        cat >>"$CONFIG" <<'EOF'
    - heading: New Delhi, India
      zone: Asia/Kolkata
      color: green
EOF
    fi
    if [[ "$name" == clocks || "$name" == twenty-four-hour ]]; then
        cat >>"$CONFIG" <<'EOF'
    - heading: New York, USA
      zone: America/New_York
      color: magenta
EOF
    fi

    export CLOCKS_DEMO_BIN="$BIN" CLOCKS_CONFIG_PATH="$CONFIG"
    export CLOCKS_DEMO_COLS="$COLS" CLOCKS_DEMO_ROWS="$ROWS" CLOCKS_DEMO_ARGS="$ARGS"
    export CLOCKS_DEMO_EXPECT="$ROOT_DIR/scripts/clocks.expect"
    if [[ "$name" == overview ]]; then
        export CLOCKS_DEMO_EXPECT="$ROOT_DIR/scripts/overview.expect"
    fi
    if [[ "$name" == twenty-four-hour ]]; then
        "$BIN" unset --t12 >/dev/null
    fi
    CAST="$CAST_DIR/$name.cast"
    RENDERED="$TMP_DIR/$name.gif"
    OUT="$ROOT_DIR/demo/$name.png"
    if [[ "$name" == overview ]]; then
        OUT="$ROOT_DIR/demo/clocks-demo.gif"
        RENDERED="$OUT"
    fi

    echo "Recording $name..."
    asciinema rec --overwrite --headless --return \
        --output-format asciicast-v2 \
        --window-size "${COLS}x${ROWS}" -i 1.4 \
        -c 'expect "$CLOCKS_DEMO_EXPECT"' "$CAST"

    agg --renderer resvg --font-family Menlo --theme github-dark --cols "$COLS" --rows "$ROWS" \
        --font-size 20 --line-height 1.0 --fps-cap 20 \
        --idle-time-limit 1.4 --last-frame-duration 3 --speed 0.85 \
        "$CAST" "$RENDERED"
    if [[ "$name" != overview ]]; then
        magick "$RENDERED" -coalesce -delete 0--2 "$OUT"
        rm -f "$ROOT_DIR/demo/$name.gif"
    else
        rm -f "$ROOT_DIR/demo/overview.gif"
    fi
    echo "Wrote $OUT"
done
